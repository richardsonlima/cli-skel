package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

var Commands map[string]cli.CommandFactory

func main() {
	log.SetOutput(ioutil.Discard)

	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("cli-skel"),
	}
	exitStatus, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(exitStatus)

	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands := map[string]cli.CommandFactory{
		"comand1": func() (cli.Command, error) {
			return &command.FirstCommand{
				Ui: ui,
			}, nil
		},
		"comand2": func() (cli.Command, error) {
			return &command.SecondCommand{
				Ui: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
	}
}
