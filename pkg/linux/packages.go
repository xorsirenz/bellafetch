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

func PkgManager() int {
	osInfo := OsRelease()
	id := osInfo["ID_LIKE"]

	if id == "" {
		id = osInfo["ID"]
	}
	if len(id) > 1 {
		id = parseID(id)
	}

	switch id {
	case "arch", "manjaro":
		pkgs := pacman()
		return pkgs
	case "debian", "linuxmint", "ubuntu":
		pkgs := dpkg()
		return pkgs
	case "void":
		pkgs := xbps()
		return pkgs
	default:
		fmt.Println("No supported package manager detected")
	}
	return 0
}

func dpkg() int {
	dpkgStatusFile := "/var/lib/dpkg/status"

	out, err := os.ReadFile(dpkgStatusFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	output := string(out)
	outputLines := strings.Split(output, "\n\n")
	lines := len(outputLines) - 1
	return lines
}

func pacman() int {
	packages := "/var/lib/pacman/local"

	entries, err := os.ReadDir(packages)
	if err != nil {
		fmt.Println("Error", err)
	}
	lines := len(entries) - 1
	return lines
}

func xbps() int {
	rootDir := "/var/db/xbps/"
	pkgdbFilePrefix := "pkgdb-"

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		fmt.Println("Error", err)
	}

	var pkgdbFile string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), pkgdbFilePrefix) {
			pkgdbFile = filepath.Join(rootDir, entry.Name())
		}
	}

	file, err := os.Open(pkgdbFile)
	if err != nil {
		fmt.Println("Error", err)
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
	return count
}
