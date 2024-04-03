package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_setupEnv(t *testing.T) {
	// save the original system env
	originGoPath := os.Getenv("GOPATH")
	originGoBin := os.Getenv("GOBIN")

	// resotre system env
	defer func() {
		os.Setenv("GOPATH", originGoPath)
		os.Setenv("GOBIN", originGoBin)
	}()

	tests := map[string]func(t *testing.T){
		"test case 1 - error with empty env": func(t *testing.T) {
			os.Unsetenv("GOPATH")
			os.Unsetenv("GOBIN")
			assert.Error(t, setupEnv())
		},
		"test case 2 - success with GOBIN": func(t *testing.T) {
			os.Unsetenv("GOPATH")
			os.Setenv("GOBIN", "/path/to/bin")
			assert.NoError(t, setupEnv())
		},
		"test case 3 - success with GOPATH": func(t *testing.T) {
			os.Setenv("GOPATH", "/path/to/go")
			os.Unsetenv("GOBIN")
			assert.NoError(t, setupEnv())
		},
		"test case 4 - success with both": func(t *testing.T) {
			os.Setenv("GOPATH", "/path/to/go")
			os.Setenv("GOBIN", "/path/to/bin")
			assert.NoError(t, setupEnv())
		},
	}

	for name, tt := range tests {
		t.Run(name, tt)
	}
}

func Test_scanVersions(t *testing.T) {
	// Create temp directory for test
	tempDir, err := os.MkdirTemp("./", "test")
	require.NoError(t, err)
	fmt.Println(tempDir)

	// clear the created directory
	defer os.RemoveAll(tempDir)

	files := []string{"go1.16.3", "gofmt", "go1.17.2", "test", "go1.18.1", "go1.19.0"}
	for _, file := range files {
		_, err := os.Create(filepath.Join(tempDir, file))
		require.NoError(t, err)
	}

	tests := map[string]func(t *testing.T){
		"test case 1 - error with invalid path": func(t *testing.T) {
			assert.Error(t, scanVersions("./path/to/invalid"))
		},
		"test case 2 - success with valid path": func(t *testing.T) {
			err := scanVersions(tempDir)
			assert.NoError(t, err)
			actual := installedVersions
			expected := []string{"go1.16.3", "go1.17.2", "go1.18.1", "go1.19.0"}
			assert.Equal(t, expected, actual)
		},
	}

	for name, tt := range tests {
		t.Run(name, tt)
	}
}
