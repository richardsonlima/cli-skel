package command

import (
	"flag"
	"strings"
)

type FirstCommand struct {
	Ui cli.Ui
}

func (c *FirstCommand) Help() string {
	helpText := `
Usage: cli-skel comand1	[options]

Options:

	-blah-blah=<value>
	-blah-blah=<value>
	-blah-blah=<value>	
`
	return strings.TrimSpace(helpText)
}

func (c *FirstCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("comand1", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (c *FirstCommand) Synopsis() string {
	return ""
}
