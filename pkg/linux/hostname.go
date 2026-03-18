package linux

import (
	"fmt"
	"strings"
)

func Hostname() string {
	hostnameFile := "/etc/hostname"

	hostnameContents, err := OpenFile(hostnameFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	hostname := string(hostnameContents)
	host := strings.TrimSuffix(hostname, "\n")
	return host
}
