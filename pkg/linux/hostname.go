package linux

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func Host() string {
	var host strings.Builder
	host.WriteString(Username())
	host.WriteString("@")
	host.WriteString(Hostname())

	return host.String()
}

func Username() string {
	user, err := user.Current()
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	return user.Username
}

func Hostname() string {
	if hostname, err := os.Hostname(); hostname != "" {
		if err != nil {
			_ = fmt.Errorf("Error: %v", err)
		}
		return hostname
	}

	hostnameFile := "/etc/hostname"
	hostname, err := os.ReadFile(hostnameFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	host := strings.TrimSuffix(string(hostname), "\n")
	return host
}
