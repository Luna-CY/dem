package echo

import "fmt"

func Infoln(message string, args ...any) {
	fmt.Printf("==> "+message+"\n", args...)
}
