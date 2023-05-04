package builtins

import (
	"fmt"
	"os"
)

// function to create one or more directories
func MakeDirectory(DirectoryNames ...string) error {
	//if statement that returns and error when no directory name is given
	if len(DirectoryNames) == 0 {
		return fmt.Errorf("%w: Please type in one or more directory name's ", ErrInvalidArgCount)
	}

	// Creates each directory with the given argument
	for _, DirectoryName:= range DirectoryNames {
		//if statement that creates directory or returns an error when directory fails to be created
		if err := os.MkdirAll(DirectoryName, 0755); err != nil {
			return fmt.Errorf("Failed to create directory %s: %v", DirectoryName, err)
		}
		// message when directory has been created
		fmt.Printf("Successfully created directory: %s\n", DirectoryName)
	}

	return nil
}
