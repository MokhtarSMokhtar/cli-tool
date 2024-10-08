package command

import (
	"flag"
	"fmt"
)

// Command represents a single CLI command.
// It holds information about the command's name, aliases, usage,
// description, flag set, and the function to run when the command is invoked.
type Command struct {
	// The name of the command (e.g., "hello")
	Name string `json:"name"`
	// A list of aliases for the command (e.g., ["hi"])
	Aliases []string `json:"aliases"`
	// A short description of how to use the command
	Usage string `json:"usage"`
	// A longer description explaining what the command does
	Description string `json:"description"`
	// The FlagSet holds all the flags associated with this command
	Flags *flag.FlagSet `json:"flags"`
	// Run is the function that is executed when the command is run
	Run func(args []string)
}

// AddNewCommand creates a new command and initializes its flag set.
// It also accepts optional aliases for the command.
// Parameters:
//
//	name - the name of the command
//	description - a brief description of what the command does
//	usage - the command's usage message
//	aliases - a variadic parameter for one or more aliases
//
// Returns:
//
//	A pointer to the created Command object.
func (c *Command) AddNewCommand(name, description string, usage string, aliases ...string) *Command {
	return &Command{
		Name:        name,
		Aliases:     aliases,
		Usage:       usage,
		Description: description,
		Flags:       flag.NewFlagSet(name, flag.ContinueOnError), // Initialize the FlagSet with name and error handling
	}
}

// Execute parses the input arguments as flags and then calls the Run function if defined.
// Parameters:
//
//	args - the arguments passed to the command (after the command itself)
func (c *Command) Execute(args []string) {
	// Parse the flags passed as arguments.
	if err := c.Flags.Parse(args); err != nil {
		// Print an error message if flag parsing fails
		fmt.Printf("Error parsing flags for command %s: %v\n", c.Name, err)
		return
	}

	// If the Run function is defined, call it with the remaining arguments.
	if c.Run != nil {
		c.Run(c.Flags.Args())
	} else {
		// If no Run function is defined, just print the usage message.
		fmt.Println(c.Usage)
	}
}
