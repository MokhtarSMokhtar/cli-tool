package main

import (
	"cli-tool/cli"
	"cli-tool/command"
	"flag"
	"fmt"
)

func main() {
	// Initialize the CLI
	cli := &cli.CLI{}

	// Define a 'hello' command
	helloCmd := &command.Command{
		Name:        "hello",
		Usage:       "hello -name <name>",
		Description: "Prints a greeting message",
		Flags:       flag.NewFlagSet("hello", flag.ExitOnError),
	}

	// Define the 'name' flag for the hello command
	var name string
	helloCmd.Flags.StringVar(&name, "name", "everyone", "Name to greet")

	// Define the command's 'Run' function
	helloCmd.Run = func(args []string) {
		// Parse the flags
		helloCmd.Flags.Parse(args)

		// Print the greeting message
		fmt.Printf("Hello, %s!\n", name)
	}

	// Register the 'hello' command to the CLI
	cli.AddCommand(helloCmd)

	// Execute the CLI with the passed arguments
	cli.Execute()
	//To run the example command
	//go run main.go hello -name John
}
