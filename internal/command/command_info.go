package command

import "github.com/gr4c2-2000/gommand/internal/proto"

type CommandInfo struct {
	ExecutableCommand string
	commandName       string
	Command           *Command
}

func CommandInfoFromGrpc(p *proto.CommandInfoResult) *CommandInfo {
	ci := &CommandInfo{
		ExecutableCommand: p.ExecutableCommand,
		commandName:       p.CommandTmp.Name,
		Command:           mapCommandTmpToCommand(p.CommandTmp),
	}
	return ci
}

func (e *CommandInfo) ToGrpc() *proto.CommandInfoResult {
	cr := &proto.CommandInfoResult{ExecutableCommand: e.ExecutableCommand,
		CommandTmp: mapCommandToCommandTmp(*e.Command, e.commandName),
	}
	return cr
}
