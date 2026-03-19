package linux

import (
	"fmt"
	"strings"

	"github.com/xorsirenz/bellafetch/pkg/utils"
)

func Hostname() string {
	hostnameFile := "/etc/hostname"

	hostnameContents, err := utils.OpenFile(hostnameFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	hostname := string(hostnameContents)
	host := strings.TrimSuffix(hostname, "\n")
	return host
}
