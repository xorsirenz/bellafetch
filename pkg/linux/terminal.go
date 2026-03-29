package linux

import (
	"os"
)

func Terminal() string {
	if term := os.Getenv("TERM_PROGRAM"); term != "" {
		return term	
	} else if term := os.Getenv("TERM"); term != "" {
		return term
	} else {
		return ""
	}
}
