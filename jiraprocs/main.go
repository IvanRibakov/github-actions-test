package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/zaleos/softin-shared/libutils/sh"
	"github.com/zaleos/softin-shared/libutils/task"
)

var (
	verbose, debug bool
)

func main() {
	// Override the version flag to use -v short flag for the verbose flag
	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}
	app := &cli.App{
		Name:        "jiraprocs",
		Version:     "2.0.0",
		Compiled:    time.Now(),
		Copyright:   "(c) 2022 Zaleos",
		Usage:       "CLI utility for generating internal JIRA reports at Zaleos",
		Description: "CLI utility for generating internal JIRA reports at Zaleos",
		Commands:    []*cli.Command{},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Destination: &verbose,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Destination: &debug,
			},
		},
		Before: func(cCtx *cli.Context) error {
			sh.Init(verbose, debug)
			task.Init(debug, false)
			return nil
		},
		EnableBashCompletion: true,
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Command %q not found.\n", command)
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}
			fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
			return nil
		},
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
	// Add sub commands here
	app.Commands = append(app.Commands, JQLCommand)
	app.Commands = append(app.Commands, PlanningCommand)
	app.Commands = append(app.Commands, GroomingCommand)
	app.Commands = append(app.Commands, VelocityCommand)

	//Run the application
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
