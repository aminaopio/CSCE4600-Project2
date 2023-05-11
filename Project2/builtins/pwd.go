package builtins 
import (
	"fmt"
	"os"
)

func Pwd(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Not working Directory: %v")
	}
	fmt.Println(wd)
	return nil
}