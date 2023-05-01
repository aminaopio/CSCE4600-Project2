package builtins

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Execute(args ...string) error {

	// prep the command
	cmdStr := strings.TrimSpace(args[0])
	cmdParts := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmdParts[0],cmdParts[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}

func Source(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := Execute(line); err != nil {
			return fmt.Errorf("failed to execute command: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return nil
}
