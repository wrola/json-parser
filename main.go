package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wojciech/json-parser/cmd"
)

func main() {
	parseCmd := flag.NewFlagSet("parse", flag.ExitOnError)
	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	formatCmd := flag.NewFlagSet("format", flag.ExitOnError)
	statsCmd := flag.NewFlagSet("stats", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "parse":
		parseCmd.Parse(os.Args[2:])
		handleParse(parseCmd.Args())
	case "validate":
		validateCmd.Parse(os.Args[2:])
		handleValidate(validateCmd.Args())
	case "format":
		formatCmd.Parse(os.Args[2:])
		handleFormat(formatCmd.Args())
	case "stats":
		statsCmd.Parse(os.Args[2:])
		handleStats(statsCmd.Args())
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("JSON Parser CLI - Learn Go fundamentals")
	fmt.Println("\nUsage:")
	fmt.Println("  json-parser <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  parse <file>      Parse and display JSON")
	fmt.Println("  validate <file>   Validate JSON syntax")
	fmt.Println("  format <file>     Format JSON with indentation")
	fmt.Println("  stats <file>      Display JSON statistics")
	fmt.Println("  help              Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  json-parser parse data.json")
	fmt.Println("  json-parser validate config.json")
	fmt.Println("  json-parser format data.json")
	fmt.Println("  json-parser stats data.json")
}

func handleParse(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: file path required")
		os.Exit(1)
	}
	if err := cmd.ParseCommand(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleValidate(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: file path required")
		os.Exit(1)
	}
	if err := cmd.ValidateCommand(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleFormat(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: file path required")
		os.Exit(1)
	}
	if err := cmd.FormatCommand(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func handleStats(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: file path required")
		os.Exit(1)
	}
	if err := cmd.StatsCommand(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
