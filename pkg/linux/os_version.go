package linux

import (
	"fmt"
	"os"
	"strings"
)

func OsRelease() map[string]string {
	OsReleaseFile := "/etc/os-release"

	contents, err := os.ReadFile(OsReleaseFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	osMap := make(map[string]string)

	for _, entry := range strings.Split(string(contents), "\n") {
		parts := strings.Split(entry, "=")
		if len(parts) == 2 {
			osMap[parts[0]] = parts[1]
		}
	}
	return osMap
}
