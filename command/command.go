package command

import (
	"flag"
	"fmt"
)

type Command struct {
	// The name of the command
	Name string `json:"name"`
	// A list of aliases for the command
	Aliases []string `json:"aliases"`
	// A short description of the usage of this command
	Usage string `json:"usage"`
	// Longer explanation of how the command works
	Description string `json:"description"`
	// the flag.FlagSet for this command
	Flags *flag.FlagSet `json:"flags"`
	// Run function to handle command execution
	Run func(args []string)
}

// Create a New Command

func (c *Command) AddNewCommand(name, description string, usage string, alies ...string) *Command {
	return &Command{
		Name:        name,
		Aliases:     alies,
		Usage:       usage,
		Description: description,
		Flags:       flag.NewFlagSet(name, flag.ContinueOnError),
	}
}

func (c *Command) Execute(args []string) {
	if err := c.Flags.Parse(args); err != nil {
		fmt.Printf("Error parsing flags for command %s: %v\n", c.Name, err)
		return
	}
	if c.Run != nil {
		c.Run(c.Flags.Args())
	} else {
		fmt.Println(c.Usage)
	}
}
