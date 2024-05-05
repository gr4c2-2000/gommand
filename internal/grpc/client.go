package grpc

import (
	"github.com/gr4c2-2000/gommand/internal/command"
	"github.com/gr4c2-2000/gommand/internal/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client proto.GommandClient
	conn   *grpc.ClientConn
}

func NewClient() (*Client, error) {
	var opts []grpc.DialOption
	var err error
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	serverAddr := "0.0.0.0:19765"
	c := Client{}
	c.conn, err = grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err

	}
	c.client = proto.NewGommandClient(c.conn)
	return &c, nil
}

func (c *Client) Info(ctx context.Context, stdIn, workDir string) (*proto.CommandInfoResult, error) {
	grpcInfo := proto.Input{StdIn: stdIn, WorkDir: workDir}
	res, err := c.client.CommandInfo(ctx, &grpcInfo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) Exec(ctx context.Context, stdin, workDir string) (*command.ExecCommandResult, error) {
	grpcCommand := proto.Command{StdIn: stdin, WorkDir: workDir}
	res, err := c.client.ExecCommand(ctx, &grpcCommand)
	if err != nil {
		return nil, err
	}
	return command.FromGrpc(res), nil
}

func (c *Client) List(ctx context.Context) ([]*proto.CommandTmp, error) {
	res, err := c.client.CommandList(ctx, new(proto.Empty))
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
