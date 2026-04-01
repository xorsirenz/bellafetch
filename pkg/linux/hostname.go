package linux

import (
	"fmt"
	"strings"
	"os"
	"os/user"
)
func Host() string {
	hostname := Hostname()
	username := Username()
	host := username + "@" + hostname
	return host
}

func Hostname() string {
	hostnameFile := "/etc/hostname"

	hostname, err := os.ReadFile(hostnameFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	host := strings.TrimSuffix(string(hostname), "\n")
	return host
}

func Username() string {
	user, err := user.Current()
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	return user.Username
}
