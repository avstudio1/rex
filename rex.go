package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Ensure the binary is invoked with at least one command
	if len(os.Args) < 2 {
		fmt.Println("Usage: rex <command> [options]")
		os.Exit(1)
	}

	// Parse the command
	command := os.Args[1]

	// Dispatch to the appropriate handler
	switch command {
	case "help":
		showHelp()
	case "topic":
		handleTopicCommand()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		showHelp()
	}
}

// showHelp displays a help menu
func showHelp() {
	fmt.Println("REX CLI Help")
	fmt.Println("Usage: rex <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  help          - Show this help menu")
	fmt.Println("  topic [name]  - Process a topic with the given name")
}

// handleTopicCommand processes the "topic" command
func handleTopicCommand() {
	// Set up a flag for the topic command
	topicCmd := flag.NewFlagSet("topic", flag.ExitOnError)
	topicName := topicCmd.String("name", "", "Name of the topic (required)")
	_ = topicCmd.Parse(os.Args[2:])

	// Validate required arguments
	if *topicName == "" {
		fmt.Println("Error: -name flag is required for the 'topic' command")
		topicCmd.Usage()
		os.Exit(1)
	}

	// Handle the topic command logic
	fmt.Printf("Processing topic: %s\n", *topicName)
}
