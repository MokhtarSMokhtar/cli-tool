package main

func main() {
	cli := &cli.CLI{}

	// Adding a 'hello' command
	helloCmd := &command.Command{
		Name:        "hello",
		Usage:       "hello -name <name>",
		Description: "Prints a greeting message",
		Flags:       flag.NewFlagSet("hello", flag.ExitOnError),
	}
	helloCmd.Run = func(args []string) {
		name := "everyone"
		fmt.Printf("Goodbye, %s!\n", name)
	}
	cli.AddCommand(helloCmd)

	// Adding a 'bye' command
	byeCmd := &command.Command{
		Name:        "bye",
		Usage:       "bye -name <name>",
		Description: "Prints a farewell message",
		Flags:       flag.NewFlagSet("bye", flag.ExitOnError),
		Run: func(args []string) {
			name := "everyone"
			fmt.Printf("Goodbye, %s!\n", name)
		},
	}
	cli.AddCommand(byeCmd)

	// Execute the CLI
	cli.Execute()
}
