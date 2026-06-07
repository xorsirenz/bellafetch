package linux

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Desktop() string {
	envVars := []string{
		"XDG_CURRENT_DESKTOP",
		"XDG_SESSION_DESKTOP",
		"DESKTOP_SESSION",
	}

	for _, v := range envVars {
		if de := os.Getenv(v); de != "" {
			return strings.ToLower(de)
		}
	}
	
	if wm := wm(); wm != "" {
		return wm
	}
	return ""
}

func wm() string {
	rootDir := "/proc"

	supportedWms := []string{
		"i3",
		"awesome",
		"bspwn",
		"dwm",
		"hyprland",
		"niri",
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
