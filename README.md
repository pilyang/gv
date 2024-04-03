# Go Version Manager

**Notice:** This is a personal toy project for study and is currently in progress. Not all features are developed yet, and there may be some bugs or unexpected results.


This project manages multiple versions of the Go programming language, installed as per the official Go documentation on [Managing Go Installations](https://go.dev/doc/manage-install). It specifically handles versions installed using the `go install` command. All managed versions are located in the `GOBIN` environment path.


## Features

### Version Check

- Display the installed Go versions with `gv -v` or `gv --versions`.
- The output is generated using `go version` or `goX.XX.X version`.

### Command Execution (In Progress)

- Run Go commands with a selected version using `gv -c (command)`.
- Display the currently installed versions in the `GOBIN` environment variable.
- Allow users to select a version to run the command.

## Planned Features

Refer to the [todo list](/todoFeatures.md) for upcoming features.


## Installation

To install the Go Version Manager, you need to have Go installed on your system. Once Go is installed, you can install the Go Version Manager using the `go install` command:

```bash
go install github.com/pilyang/gv@latest
```

This command will install the latest version of the Go Version Manager.


## User Guide

### Version Check

To check the installed Go versions, you can use the following commands:

```bash
gv -v
```
or
```bash
gv --version
```

These commands will display the Go versions that are installed on your system. The first line of the output shows the Go version installed in the `GOROOT` path, which is typically installed from the official installer package or Homebrew (for macOS). The following lines show all the Go versions installed in the `GOBIN` path.

Here's an example of the output you might see:


```bash
$ gv -v
GOROOT Installed version
go version go1.22.1 darwin/arm64

GOBIN Installed versions
go version go1.18.1 darwin/arm64
```

