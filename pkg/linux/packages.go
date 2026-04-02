package linux

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseID(id string) string {
	ids := strings.Fields(id)
	firstID := ids[0]
	cleanedID := strings.ReplaceAll(firstID, "\"", "")
	id = cleanedID
	return id
}

func PkgManager(osMap map[string]string) string {
	id := osMap["ID_LIKE"]

	if id == "" {
		id = osMap["ID"]
	}
	if len(id) > 1 {
		id = parseID(id)
	}

	flatpaks := flatpak()

	switch id {
	case "arch", "manjaro":
		pkgs := pacman()
		return fmt.Sprintf("%s %s", pkgs, flatpaks)
	case "debian", "linuxmint", "ubuntu":
		pkgs := dpkg()
		return fmt.Sprintf("%s %s", pkgs, flatpaks)
	case "void":
		pkgs := xbps()
		return fmt.Sprintf("%s %s", pkgs, flatpaks)
	case "rhel":
		return fmt.Sprintf("%s", flatpaks)
	case "nixos":
		pkgs := nixos()
		return fmt.Sprintf("%s %s", pkgs, flatpaks)
	default:
		fmt.Println("No supported package manager detected")
	}
	return ""
}

func flatpak() string {
	appTotal := flatpakApps()
	runtimeTotal := flatpakRuntimes()
	flatpakTotal := appTotal + runtimeTotal
	if flatpakTotal != 0 {
		return fmt.Sprintf("%d (flatpak)", flatpakTotal)
	}
	return ""
}

func flatpakApps() int {
	appDir := "/var/lib/flatpak/app/"

	entries, err := os.ReadDir(appDir)
	if err != nil {
		return 0
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		appName := entry.Name()
		currentPath := filepath.Join(appDir, appName, "current")
		if strings.HasSuffix(currentPath, "current") {
			count++
		}
	}
	return count
}

func flatpakRuntimes() int {
	runtimeDir := "/var/lib/flatpak/runtime/"

	entries, err := os.ReadDir(runtimeDir)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && !strings.HasSuffix(entry.Name(), ".Locale") || !strings.HasSuffix(entry.Name(), ".Debug") {
			count++
		}
	}
	return count
}

func nixos() string {
	rootDir := "/run/current-system/sw/bin"

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		_ = fmt.Errorf("Error %v", err)
	}
	lines := len(entries) - 1
	return fmt.Sprintf("%d (nixos)", lines)
}

func dpkg() string {
	dpkgStatusFile := "/var/lib/dpkg/status"

	out, err := os.ReadFile(dpkgStatusFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	output := string(out)
	outputLines := strings.Split(output, "\n\n")
	lines := len(outputLines) - 1
	return fmt.Sprintf("%d (dpkg)", lines)
}

func pacman() string {
	rootDir := "/var/lib/pacman/local"

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		_ = fmt.Errorf("Error %v", err)
	}
	lines := len(entries) - 1
	return fmt.Sprintf("%d (pacman)", lines)
}

func xbps() string {
	rootDir := "/var/db/xbps/"
	pkgdbFilePrefix := "pkgdb-"

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	var pkgdbFile string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), pkgdbFilePrefix) {
			pkgdbFile = filepath.Join(rootDir, entry.Name())
		}
	}

	file, err := os.Open(pkgdbFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	installed := "<string>installed</string>"
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, installed)
	}
	return fmt.Sprintf("%d (xbps)", count)
}
