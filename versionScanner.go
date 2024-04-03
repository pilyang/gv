package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

var (
	gobin             string
	gopath            string
	installedVersions []string

	reg = regexp.MustCompile(`go\d+\.\d+\.\d+`)
)

func setupEnv() error {
	gobin = os.Getenv("GOBIN")
	gopath = os.Getenv("GOPATH")

	if gobin == "" && gopath == "" {
		return errors.New("GOPATH and GOBIN are not set")
	}

	if gobin == "" {
		gobin = gopath + "/bin"
	}
	return nil
}

func scanVersions(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	installedVersions = make([]string, 0)
	for _, file := range files {
		if reg.MatchString(file.Name()) {
			installedVersions = append(installedVersions, file.Name())
		}
	}
	return nil
}

func init() {
	err := setupEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = scanVersions(gobin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func InstalledVersions() []string {
	return installedVersions
}
