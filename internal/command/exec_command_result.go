package command

import "github.com/gr4c2-2000/gommand/internal/proto"

type ExecCommandResult struct {
	Stdout, Stderr string
	Err            error
	Command        string
}

func FromGrpc(p *proto.CommandResult) *ExecCommandResult {
	return &ExecCommandResult{Stdout: p.Stdout, Stderr: p.StdErr, Command: p.Command}
}

func (e *ExecCommandResult) ToGrpc() *proto.CommandResult {
	stdErr := e.Stderr
	if e.Err != nil {
		stdErr = stdErr + "\n" + e.Err.Error()
	}
	cr := proto.CommandResult{Stdout: e.Stdout, StdErr: stdErr, Command: e.Command}
	return &cr
}
