package linux

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Wm() string {
	rootDir := "/proc"

	supportedWms := []string{
		"i3",
		"awesome",
		"bspwn",
		"dwm",
		"hyprland",
		"openbox",
		"river",
		"sway",
	}

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		return ""
	}

	for _, entry := range entries {
		pidName := entry.Name()
		if _, err := strconv.Atoi(pidName); err != nil {
			continue
		}

		commPath := filepath.Join(rootDir, pidName, "comm")
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
