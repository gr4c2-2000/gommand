package grpc

import (
	"context"
	"errors"
	"net"

	"github.com/gr4c2-2000/gommand/internal/command"
	"github.com/gr4c2-2000/gommand/internal/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGommandServer
	execService   *command.ExecService
	manageService *command.ManageService
}

func NewServer(execs *command.ExecService, manags *command.ManageService) error {
	tcpl, err := net.Listen("tcp", ":19765")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGommandServer(grpcServer, &Server{execService: execs, manageService: manags})

	if err := grpcServer.Serve(tcpl); err != nil {
		return err
	}

	return nil
}

func (s *Server) ExecCommand(ctx context.Context, req *proto.Command) (*proto.CommandResult, error) {
	var err error
	if req.Command == "" {
		err := errors.New("gommand : incorrect command : " + req.Command)
		return nil, err
	}
	execCommandResult, err := s.execService.ExecServiceStrategy(req.Command, req.WorkDir)
	if err != nil {
		return nil, err
	}
	return execCommandResult.ToGrpc(), err
}

func (s *Server) CommandList(ctx context.Context, in *proto.Empty) (*proto.CommandListResult, error) {
	res := command.StorageToCommandListResult(s.manageService.GetCommandTemplates())
	return res, nil
}
