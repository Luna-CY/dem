package echo

import "fmt"

func Info(message string, args ...interface{}) {
	fmt.Printf("==> "+message+"\n", args...)
}
