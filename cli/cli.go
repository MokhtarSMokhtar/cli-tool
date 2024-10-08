package cli

import (
	"cli-tool/command"
	"fmt"
	"os"
)

// Registering Commands
type CLI struct {
	Commands []*command.Command
}

// AddCommand to the CLI
func (cli *CLI) AddCommand(cmd *command.Command) {
	cli.Commands = append(cli.Commands, cmd)
}

// Find and Execute a command based on input arguments

func (cli *CLI) Execute() {
	if len(os.Args) < 2 {
		fmt.Println("You must pass a command")
		return
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	for _, cmd := range cli.Commands {
		if cmd.Name == commandName || contains(cmd.Aliases, commandName) {
			cmd.Execute(args)
			return
		}
	}

	fmt.Printf("Unknown command: %s\n", commandName)
}

// Helper function to check for aliases
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
