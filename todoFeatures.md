# Project TODO List

This project is focused on creating a version manager for the Go programming language. The features listed here are designed to manage the installed versions as per the [official Go documentation](https://go.dev/doc/manage-install).


## MVP Features

- [ ] **Version Check**
  - Implement the `gv -versions` command
  - Scan and display the installed Go versions
  - Output should be generated using `go version` or `goX.XX.X version`

- [ ] **Command Execution**
  - Implement the `gv -c (command)` feature
  - Allow users to run Go commands with a selected version
  - Display the currently installed versions in the `GOBIN` environment variable
  - Allow users to select a version to run the command


## Optional Features

These features are not essential for the initial release but could enhance the functionality of the Go version manager:

- [ ] **Version Installation**
  - Implement the `gv -i x.x.x` command to install new Go versions

- [ ] **Version Removal**
  - Implement a feature to remove existing Go versions in the `GOBIN` environment variable

- [ ] **Version Tagging**
  - Implement a feature to add tags to installed versions
  - Allow users to select or search versions with tags


## Further Considerations

This section is for any additional thoughts, ideas, or considerations that may affect the project in the future.

I'm not familiar with Golang, so I'm not sure if these features (including the above MVP and Optional features) are helpful or not.

- Uncertain if it will be possible to use a specific version with these commands
