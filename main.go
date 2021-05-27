package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

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
}

// func main() {
// 	c := cli.NewCLI("app", "1.0.0")
// 	c.Args = os.Args[1:]
// 	c.Commands = map[string]cli.CommandFactory{
// 		"foo": fooCommandFactory,
// 		"bar": barCommandFactory,
// 	}

// 	exitStatus, err := c.Run()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	os.Exit(exitStatus)
// }
