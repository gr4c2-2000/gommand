package daemon

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

var stdlog, errlog *log.Logger

type Process interface {
	Run() error
	Interrupt(os.Signal)
}

type Daemon struct {
	daemon.Daemon
	Process Process
}

func (dae *Daemon) Manage() (string, error) {

	usage := "Usage: gommandd install | remove | start | stop | status"
	// If received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return dae.Install()
		case "remove":
			return dae.Remove()
		case "start":
			return dae.Start()
		case "stop":
			// No need to explicitly stop cron since job will be killed
			return dae.Stop()
		case "status":
			return dae.Status()
		default:
			return usage, nil
		}
	}
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	err := dae.Process.Run()
	if err != nil {
		errlog.Println("Cannot run daemon err: %w", err)
	}
	// Waiting for interrupt by system signal
	killSignal := <-interrupt
	dae.Process.Interrupt(killSignal)
	stdlog.Println("Got signal:", killSignal)
	return "Service exited", nil
}

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func NewDaemonService(name, desc string, proc Process) (*Daemon, error) {
	srv, err := daemon.New(name, desc, daemon.SystemDaemon)
	if err != nil {
		return nil, err
	}
	return &Daemon{srv, proc}, nil
}

func (dae *Daemon) Run() {
	status, err := dae.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
