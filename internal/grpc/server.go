package grpc

import (
	"context"
	"errors"
	"net"
	"os"

	"github.com/gr4c2-2000/gommand/internal/command"
	"github.com/gr4c2-2000/gommand/internal/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGommandServer
	execService   *command.ExecService
	manageService *command.ManageService
	context       context.Context
	grpcServ      *grpc.Server
	mainErrChan   chan error
}

type DaemonServerWrapper struct {
	server *Server
	execs  *command.ExecService
	manags *command.ManageService
}

func NewDaemonServerWrapper(execs *command.ExecService, manags *command.ManageService) *DaemonServerWrapper {
	return &DaemonServerWrapper{execs: execs, manags: manags}

}

func (dsw *DaemonServerWrapper) Run() error {
	var err error
	if dsw.server != nil {
		dsw.server.Close()
	}
	dsw.server, err = NewServer(context.Background(), dsw.execs, dsw.manags)
	return err
}

func (dsw *DaemonServerWrapper) Interrupt(os.Signal) {
	if dsw.server != nil {
		dsw.server.Close()
	}
}

func NewServer(ctx context.Context, execs *command.ExecService, manags *command.ManageService) (*Server, error) {
	tcpl, err := net.Listen("tcp", ":19765")
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	server := &Server{execService: execs,
		manageService: manags,
		context:       ctx,
		grpcServ:      grpcServer,
		mainErrChan:   make(chan error),
	}
	proto.RegisterGommandServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(tcpl); err != nil {
			server.mainErrChan <- err
		}
	}()

	return server, nil
}

func (s *Server) Wait() error {
	defer close(s.mainErrChan)
	return <-s.mainErrChan

}
func (s *Server) Close() {
	s.grpcServ.GracefulStop()
	close(s.mainErrChan)
}

func (s *Server) Stop() {
	s.grpcServ.Stop()
	close(s.mainErrChan)
}

func (s *Server) CommandInfo(ctx context.Context, req *proto.Input) (*proto.CommandInfoResult, error) {
	var err error
	if req.StdIn == "" {
		err := errors.New("gommand : incorrect command : " + req.StdIn)
		return nil, err
	}
	commandInfo, err := s.execService.GetCommandInfo(req.StdIn)
	if err != nil {
		return nil, err
	}
	return commandInfo.ToGrpc(), err
}

func (s *Server) ExecCommand(ctx context.Context, req *proto.Command) (*proto.CommandResult, error) {
	var err error
	if req.StdIn == "" {
		err := errors.New("gommand : incorrect command : " + req.StdIn)
		return nil, err
	}
	execCommandResult, err := s.execService.ExecServiceStrategy(req.StdIn, req.WorkDir)
	if err != nil {
		return nil, err
	}
	return execCommandResult.ToGrpc(), err
}

func (s *Server) CommandList(ctx context.Context, in *proto.Empty) (*proto.CommandListResult, error) {
	res := command.StorageToCommandListResult(s.manageService.GetCommandTemplates())
	return res, nil
}
