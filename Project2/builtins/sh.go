package builtins

import (
	"fmt"
	"os/exec"
)

// function to execute shell commands
func BourneShell(command string) error {
	//if statement that returns an error when no command is given
	if command == "" {
		return fmt.Errorf("%w: Please type in a shell command", ErrInvalidArgCount)
	}

	// Runs the shell command with the given argument
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to run command %s: %v", command, err)
	}

	return nil
}
