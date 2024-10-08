package cli

import (
	"fmt"
	"github.com/MokhtarSMokhtar/cli-tool/command"
	"os"
)

// CLI represents the main command-line interface that holds and manages all the registered commands.
// more about struct https://github.com/MokhtarSMokhtar/Golan-Concepts/blob/main/pkg/structs/structs.go
type CLI struct {
	// Commands holds all the registered commands for the CLI.
	Commands []*command.Command
}

// AddCommand adds a new command to the CLI's list of commands.
// Parameters:
//
//	cmd - the command object to be added.
func (cli *CLI) AddCommand(cmd *command.Command) {
	cli.Commands = append(cli.Commands, cmd)
}

// Execute processes the input arguments, finds the appropriate command,
// and executes it. If no valid command is found, it prints an error message.
// It also handles the scenario where the user hasn't provided a command.
func (cli *CLI) Execute() {
	// Check if there are enough arguments (i.e., the command and its flags).
	if len(os.Args) < 2 {
		fmt.Println("You must pass a command")
		return
	}

	// The first argument is the command name.
	commandName := os.Args[1]
	args := os.Args[2:]

	// Loop through the registered commands and check if the command name or alias matches.
	for _, cmd := range cli.Commands {
		if cmd.Name == commandName || contains(cmd.Aliases, commandName) {
			cmd.Execute(args)
			return
		}
	}

	// If no match is found, print an error message.
	fmt.Printf("Unknown command: %s\n", commandName)
}

// Helper function that checks whether a given string exists in a list of aliases.
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
