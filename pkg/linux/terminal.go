package linux

import (
	"os"
)

func Terminal() string {
	if term := os.Getenv("TERM_PROGRAM"); term != "" {
		return term
	}
	return ""
}
