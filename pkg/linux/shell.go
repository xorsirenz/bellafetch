package linux

import (
	"os"
)

func Shell() string {
	if shell := os.Getenv("SHELL"); shell != "" {
		return shell[1:]
	}
	return ""
}
