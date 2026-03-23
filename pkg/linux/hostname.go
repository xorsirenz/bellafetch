package linux

import (
	"fmt"
	"strings"
	"os/user"

	"github.com/xorsirenz/bellafetch/internal/utils"
)
func Host() string {
	hostname := Hostname()
	username := Username()
	host := hostname + "@" + username
	return host
}

func Hostname() string {
	hostname := "/etc/hostname"

	hostnameContents, err := utils.OpenFile(hostname)
	if err != nil {
		fmt.Println("Error:", err)
	}

	host := strings.TrimSuffix(hostnameContents, "\n")
	return host
}

func Username() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user.Username
}
