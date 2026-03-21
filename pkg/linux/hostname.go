package linux

import (
	"fmt"
	"strings"

	"github.com/xorsirenz/bellafetch/pkg/utils"
)

func Hostname() string {
	hostname := "/etc/hostname"

	hostnameContents, err := utils.OpenFile(hostname)
	if err != nil {
		fmt.Println("Error:", err)
	}

	host := strings.TrimSuffix(hostnameContents, "\n")
	return host
}
