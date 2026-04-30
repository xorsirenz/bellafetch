package utils

import (
	"fmt"
)

func FetchColorBlock(config Config) []string {
	if config.ColorBlocks == false {
		return nil
	}
	var lines []string
	var line string

	for i := 0; i < 16; i++ {
		line += fmt.Sprintf("\x1b[48;5;%dm    \x1b[0m", i)

		if i == 7 {
			lines = append(lines, line)
			line = ""
		}
	}

	if line != "" {
		lines = append(lines, line)
	}

	return lines
}
