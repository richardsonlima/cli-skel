package command

import (
	"flag"
	"strings"
)

type SecondCommand struct {
	Ui cli.Ui
}

func (c *SecondCommand) Help() string {
	helpText := `
Usage: cli-skel comand2	[options]

Options:

	-blah-blah=<value>
	-blah-blah=<value>
	-blah-blah=<value>	
`
	return strings.TrimSpace(helpText)
}

func (c *SecondCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("comand2", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (c *SecondCommand) Synopsis() string {
	return ""
}
