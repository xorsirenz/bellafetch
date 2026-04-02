package linux

import (
	"fmt"
	"strings"
	"os"
	"os/user"
)
func Host() string {
	var host strings.Builder
	host.WriteString(Hostname())
	host.WriteString("@")
	host.WriteString(Username())

	return host.String()
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
