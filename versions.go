package main

import (
	"fmt"
	"os/exec"
)

const version = "version"

// ShowVersions prints the installed Go versions.
// It first prints the version of Go installed in GOROOT,
// then it prints the versions of Go installed in GOBIN.
func ShowVersions() {
	// Print the version of Go installed in GOROOT
	fmt.Println("GOROOT Installed version")
	executeAndWriteOutput("go")

	// Get the versions of Go installed in GOBIN
	versions := InstalledVersions()

	// Print the versions of Go installed in GOBIN
	fmt.Println("\nGOBIN Installed versions")
	for _, v := range versions {
		executeAndWriteOutput(v)
	}
}

func executeAndWriteOutput(command string) {
	cmd := exec.Command(command, version)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		return
	}
	fmt.Print(string(output))
}
