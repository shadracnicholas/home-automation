package main

import (
	"os"

	"github.com/shadracnicholas/home-automation/tools/bolt/cmd"
	"github.com/shadracnicholas/home-automation/tools/deploy/pkg/output"
)

// BuildDirectory is injected at compile time
var BuildDirectory string

func main() {
	if cwd, err := os.Getwd(); err != nil {
		output.Fatal("Failed to get pwd: %v", err)
	} else if cwd != BuildDirectory {
		output.Fatal("Must be run from home-automation root: %s\n", BuildDirectory)
	}

	cmd.Execute()
}
