package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gr4c2-2000/gommand/internal/command"
	"github.com/gr4c2-2000/gommand/internal/grpc"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "gmd",
	Long: "Root command",
}

type App struct {
	grpcClient *grpc.Client
}

func main() {
	app := AppInit()
	RootCmd.AddCommand(app.GenerateCommands()...)
	defer app.grpcClient.Close()

	RootCmd.AddCommand(app.GetDefault())

	cmd, _, err := RootCmd.Find(os.Args[1:])
	//TODO : zmienić na walidację czy os.Args[1:][0] jest na liście komend z templatek
	if (err != nil || cmd == nil) && os.Args[1:][0] != "completion" && os.Args[1:][0] != "__complete" {
		args := append([]string{"default"}, os.Args...)
		RootCmd.SetArgs(args)
	}

	if err := RootCmd.Execute(); err != nil {
		log.Default().Printf("error exec: %v", err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func AppInit() *App {
	client, err := grpc.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return &App{grpcClient: client}
}

func (a *App) GetDefault() *cobra.Command {
	return &cobra.Command{
		Use:    "default",
		Long:   "default command runner",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			dir, _ := os.Getwd()
			a.exec(args[1:], dir)
		},
	}
}

func (a *App) exec(args []string, workDir string) {
	commandInfoProto, err := a.grpcClient.Info(context.Background(), strings.Join(args, " "), workDir)
	if err != nil {
		//TODO : Change to stderr
		log.Fatal(err)
	}
	commandInfo := command.CommandInfoFromGrpc(commandInfoProto)
	a.ExecCommand(commandInfo, args, workDir)
}

func (a *App) ExecCommand(commandInfo *command.CommandInfo, args []string, workDir string) {
	if commandInfo.Command.Sync {
		a.execSync(commandInfo, args, workDir)
		return
	}
	a.execAsync(args, workDir)
}

func (a *App) execSync(commandInfo *command.CommandInfo, args []string, workDir string) {
	execArgs := []string{}
	if len(args) > 0 {
		execArgs = args[1:]
	}
	err := command.ExecSync(commandInfo.Command.Shell, commandInfo.ExecutableCommand, execArgs, workDir)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) execAsync(args []string, workDir string) {
	execCommandResult, err := a.grpcClient.Exec(context.Background(), strings.Join(args, " "), workDir)
	if err != nil {
		//TODO : Change to stderr
		log.Fatal(err)

	}
	if execCommandResult.Stderr != "" {
		fmt.Println(execCommandResult.Command)
		fmt.Println(execCommandResult.Stdout)
		log.Fatal(execCommandResult.Stderr)
	}

	fmt.Println(execCommandResult.Command)
	fmt.Println(execCommandResult.Stdout)
	fmt.Print(execCommandResult.Stderr)

}

func (a *App) GenerateCommands() []*cobra.Command {
	ret := []*cobra.Command{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := a.grpcClient.List(ctx)
	if err != nil {
		return nil
	}
	for _, item := range res {
		ret = append(ret, &cobra.Command{
			Use:   item.Name,
			Short: item.Name + " Alias for : " + item.ExecTmp + " ",
			Run: func(cmd *cobra.Command, args []string) {
				dir, _ := os.Getwd()
				args = append([]string{item.Name}, args...)
				a.exec(args, dir)
			},
		})
	}
	return ret
}
