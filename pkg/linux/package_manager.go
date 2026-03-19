package linux

import (
	"fmt"
	"os"
	"strings"

	"github.com/xorsirenz/bellafetch/pkg/utils"
)

func PkgManager() int {
	osInfo := OsRelease()
	id := osInfo["ID_LIKE"]

	if id == "" {
		id = osInfo["ID"]
	}

	switch id {
	case "arch", "manjaro":
		pkgs := pacman()
		return pkgs
	case "debian", "linuxmint", "ubuntu":
		pkgs := dpkg()
		return pkgs
	default:
		fmt.Println("No supported package manager detected")
	}
	return 0
}

func dpkg() int {
	dpkgStatusFile := "/var/lib/dpkg/status"

	out, err := utils.OpenFile(dpkgStatusFile)
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

