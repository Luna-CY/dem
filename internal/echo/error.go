package echo

import (
	"fmt"
	"os"
)

func Error(message string, args ...interface{}) error {
	_, _ = fmt.Fprintf(os.Stderr, "==> "+message+"\n", args...)

	return fmt.Errorf(message, args...)
}
