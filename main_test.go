package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_extractArguments(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    []argument
		wantErr bool
	}{
		{
			name: "no arguments",
			args: []string{},
			want: []argument{},
		},
		{
			name: "single command with no args",
			args: []string{"-c"},
			want: []argument{{command: "-c"}},
		},
		{
			name: "single command with one arg",
			args: []string{"--command", "arg1"},
			want: []argument{{command: "--command", args: []string{"arg1"}}},
		},
		{
			name: "single command with multiple args",
			args: []string{"-c", "arg1", "arg2"},
			want: []argument{{command: "-c", args: []string{"arg1", "arg2"}}},
		},
		{
			name: "multiple commands with multiple args",
			args: []string{"-c1", "arg1", "arg2", "--command2", "arg3", "arg4"},
			want: []argument{
				{command: "-c1", args: []string{"arg1", "arg2"}},
				{command: "--command2", args: []string{"arg3", "arg4"}},
			},
		},
		{
			name:    "arg without command",
			args:    []string{"arg1"},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := extractArguments(tc.args)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.want, actual)
			}
		})
	}
}
