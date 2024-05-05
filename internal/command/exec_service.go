package command

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

type ExecService struct {
	Storage *Storage
}

func NewExecService(s *Storage) *ExecService {
	es := ExecService{Storage: s}
	return &es
}

func (es *ExecService) GetCommandInfo(stdin string) (*CommandInfo, error) {
	ci := new(CommandInfo)
	var err error
	ci.Command, err = es.Storage.FindCommand(stdin)
	if err != nil {
		return nil, err
	}
	ci.ExecutableCommand, err = es.recursiveWrapping(ci.Command)
	if err != nil {
		return nil, err
	}

	return ci, nil

}

func (es *ExecService) ExecServiceStrategy(stdin string, workDir string) (*ExecCommandResult, error) {
	var cmdstr string
	WrapDeffer := func() {}
	ctx := context.Background()

	command, err := es.Storage.FindCommand(stdin)
	if err != nil {
		return nil, err
	}

	if command.WorkDir != "" {
		workDir = command.WorkDir
	}

	cmdstr, err = es.recursiveWrapping(command)
	if err != nil {
		return nil, err
	}
	if command.Timeout > 0 {
		ctx, WrapDeffer = context.WithTimeout(ctx, command.Timeout)
	}

	execFunc := func() *ExecCommandResult {
		defer WrapDeffer()
		return ShellExec(ctx, command.Shell, cmdstr, workDir)
	}
	if !command.Sync {
		go execFunc()
		return &ExecCommandResult{Stdout: "Task dispatched", Command: cmdstr}, nil
	}

	return execFunc(), nil
}

func ShellExec(ctx context.Context, shell, c string, workDir string) *ExecCommandResult {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, shell, "-c", c)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if workDir != "" {
		cmd.Dir = workDir
	}

	err := cmd.Run()

	return &ExecCommandResult{Stdout: stdout.String(), Stderr: stderr.String(), Err: err, Command: c}
}

func (es *ExecService) recursiveWrapping(command *Command) (string, error) {
	var err error
	var HasWrapper = true
	recursiveCommand := command

	cmdstr, err := command.GetCommand()
	if err != nil {
		return "", err
	}

	for HasWrapper {
		if strings.TrimSpace(recursiveCommand.WrapperName) == "" {
			HasWrapper = false
			break
		}
		recursiveCommand, err = es.Storage.FindCommand(recursiveCommand.WrapperName + "-\"" + cmdstr + "\"")
		if err != nil {
			return "", err
		}
		cmdstr, err = recursiveCommand.GetCommand()
		if err != nil {
			return "", err
		}
	}
	return cmdstr, err
}
