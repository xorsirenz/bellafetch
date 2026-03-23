package linux

import (
	"strings"
	"fmt"

	"github.com/xorsirenz/bellafetch/internal/utils"
)
func OsRelease() map[string]string {
	OsReleaseFile := "/etc/os-release"

	contents, err := utils.OpenFile(OsReleaseFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	entries := strings.Split(contents, "\n")
	osMap := make(map[string]string)

	for _, entry := range entries {
		parts := strings.Split(entry, "=")
		if len(parts) == 2 {
			osMap[parts[0]] = parts[1]
		}
	}
	return osMap
}

