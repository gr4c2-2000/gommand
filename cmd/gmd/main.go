package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

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

	// f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// log.Default().SetOutput(f)
	// log.Default().Printf("args:%v", os.Args)
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
	execCommandResult, err := a.grpcClient.Exec(context.Background(), strings.Join(args, " "), workDir)
	if err != nil {
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
