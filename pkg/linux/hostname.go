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
	host := hostname + "@" + username
	return host
}

func Hostname() string {
	hostnameFile := "/etc/hostname"

	hostname, err := os.ReadFile(hostnameFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	host := strings.TrimSuffix(string(hostname), "\n")
	return host
}

func Username() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user.Username
}
