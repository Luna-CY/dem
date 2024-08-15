package echo

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Errorln(message string, ps bool, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "==> "+message+"\n", args...)

	if ps {
		debug.PrintStack()
	}
}
