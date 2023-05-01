package builtins

import (
	"fmt"
	"io"
)

func Help(w io.Writer, commands map[string]string) error {
	_, err := fmt.Fprintf(w, "Available commands:\n")
	if err != nil {
		return err
	}
	for name, desc := range commands {
		_, err := fmt.Fprintf(w, "  %s: %s\n", name, desc)
		if err != nil {
			return err
		}
	}
	return nil
}
