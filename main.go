package main

import (
	"errors"
	"fmt"
	"os"
)

type argument struct {
	command string
	args    []string
}

// extracts the command and arguments from the command line arguments.
// Command starts with '-' or '--' and arguments are the rest of the arguments.
// there could be multiple commands and arguments.
// Example:
// go run main.go -v --version -h --help
// will return
//
//	[]argument{
//		{command: "-v", args: []string{}},
//		{command: "--version", args: []string{}},
//		{command: "-h", args: []string{}},
//		{command: "--help", args: []string{}},
//	}
func extractArguments(args []string) ([]argument, error) {
	arguments := make([]argument, 0)
	var current argument

	for _, arg := range args {
		if arg[0] == '-' {
			if current.command != "" {
				arguments = append(arguments, current)
			}
			current = argument{command: arg}
			continue
		}
		if current.command == "" {
			return nil, errors.New("argument " + arg + " without command")
		}
		current.args = append(current.args, arg)
	}

	if current.command != "" {
		arguments = append(arguments, current)
	}

	return arguments, nil
}

func main() {
	// os.Args[1:] is used to remove the program name from the arguments
	arguments, err := extractArguments(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, argument := range arguments {
		switch argument.command {
		case "-v", "--version":
			ShowVersions(argument.args...)
		}
	}
}
