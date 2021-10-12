package utils

import (
	"fmt"
	"os"
)

func FatalError(err error) {
	if err == nil {
		return
	}
	fmt.Printf("Fatal error: %s\n", err)
	os.Exit(1)
}
