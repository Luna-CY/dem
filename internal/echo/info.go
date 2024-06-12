package echo

import "fmt"

func Info(message string, args ...interface{}) error {
	fmt.Printf("==> "+message+"\n", args...)

	return nil
}
