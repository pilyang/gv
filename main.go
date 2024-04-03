package main

import "os"

func main() {
	args := os.Args

	// if first argument is "-v" or "--version" show versions
	if len(args) > 1 && (args[1] == "-v" || args[1] == "--version") {
		ShowVersions()
	}
}
