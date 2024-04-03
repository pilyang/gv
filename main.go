package main

import "fmt"

func main() {
	versions := InstalledVersions()
	for _, version := range versions {
		fmt.Println(version)
	}
}
