package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/richardsonlima/cli-skel/command"
)

var Commands map[string]cli.CommandFactory

func init() {
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
