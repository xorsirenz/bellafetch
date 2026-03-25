package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Wm() string {
	rootDir := "/proc"

	supportedWms := []string {
		"i3",
		"awesome",
		"bspwn",
		"dwm",
		"hyprland",
		"openbox",
		"river",
		"sway",
	}

	numericDirRegex := regexp.MustCompile(`^\d`)

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Printf("Error readinging file: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir()  || !numericDirRegex.MatchString(entry.Name()){
			continue
		}

		pidname := entry.Name()
		commPath := filepath.Join(rootDir, pidname, "comm")
		data, err := os.ReadFile(commPath)
		if err != nil {
			continue
		}

		commFile := strings.TrimSpace(string(data))
		for _, wm := range supportedWms {
			if commFile == wm {
				return wm
			}
		}
	}
	return ""
}
