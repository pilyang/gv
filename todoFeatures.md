# Project TODO List

All features here will handle the versions installed as per this [document](https://go.dev/doc/manage-install).

## MVP Features

- [ ] Check the installed versions
  - Use the `goversions` command
  - Scan the installed versions
  - Display the installed versions info using `go version` or `goX.XX.X version`

## Optional Features

These are non-essential features for the MVP but would be nice to have. They may or may not be developed in the future.

- [ ] Use arguments for the command
- [ ] Install a new version like `goversions -i x.x.x`
- [ ] Remove an existing version
- [ ] Change the default Go version registered in GOROOT
  - The location of the default version may differ from the versions installed by `go install`. This should be handled correctly.
- [ ] Manage the installed versions with an alias for each version
  - For changing the default version
  - Verify the version in use (may be closer to a tag)

## Further Considerations

This section is for any additional thoughts, ideas, or considerations that may affect the project in the future.

I'm not familiar with Golang, so I'm not sure if these features (including the above MVP and Optional features) are helpful or not.

- Uncertain if it will be possible to use a specific version with these commands
