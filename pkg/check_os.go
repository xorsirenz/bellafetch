package pkg

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xorsirenz/bellafetch/internal/utils"
	"github.com/xorsirenz/bellafetch/pkg/linux"
)

func CheckOS() utils.Data {
	goos := runtime.GOOS
	switch goos {
	case "linux":
		return linux.GetLinuxData()
	case "freebsd", "netbsd", "openbsd", "dragonfly":
		fmt.Println("Error: Bellafetch is not capitable with any BSD derivatives right now.")
		os.Exit(-1)
	case "darwin":
		fmt.Println("Error: Bellafetch is not capitable with Darwin/Mac OSX right now..")
		os.Exit(-1)
	case "windows":
		fmt.Println("Error: Bellafetch is not capitable with Windows right now.")
		os.Exit(-1)
	default:
		fmt.Println("Error: Bellafetch cannot detect OS target.")
		os.Exit(-1)
	}
	return utils.Data{}
}
