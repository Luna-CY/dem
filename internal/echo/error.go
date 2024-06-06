package echo

import "fmt"

func Error(message string, args ...interface{}) error {
	fmt.Printf(message+"\n", args...)

	return nil
}
