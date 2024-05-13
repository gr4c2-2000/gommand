package command

import (
	"context"
	"os"
	"os/exec"
)

// type ExecCommand struct {
// 	ctx     context.Context
// 	shell   string
// 	command string
// 	args    []string
// 	workDir string
// }

func ExecSync(shell, command string, args []string, workDir string) error {
	return execSync(context.Background(), shell, command, args, workDir)
}

func execSync(ctx context.Context, shell, command string, args []string, workDir string) error {
	// Execute the command
	args2 := append([]string{"-c"}, command)
	args2 = append(args2, args...)
	cmd := exec.CommandContext(ctx, shell, args2...)

	// Set the standard input/output of the command to the current process's
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if workDir != "" {
		cmd.Dir = workDir
	}

	// Start the command
	err := cmd.Start()
	if err != nil {
		return err
	}

	// Wait for the command to finish
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
