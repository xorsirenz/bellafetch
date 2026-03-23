package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xorsirenz/bellafetch/internal/utils"
	"github.com/xorsirenz/bellafetch/pkg/linux"
)

var version string

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func checkOS() utils.Data {
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

func main() {
	data := checkOS()
	clearScreen()

	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Error: Cannot load config file.")
	}

	contextMap := map[string]string{
		"Host":       "  host    ::",
		"PrettyName": "  os      ::",
		"Kernel":     "  ver     ::",
		"Uptime":     "  uptime  ::",
		"Packages":   "  pkgs    ::",
		"WM":         "  wm      ::",
		"Cpu":        "  cpu     ::",
		"Gpu":        "  gpu     ::",
		"DiskSpace":  "  storage ::",
		"Memory":     " memory  ::",
	}

	const banner = `
	 bellafetch
    [github : xorsirenz]
	`

	fmt.Println(banner)
	utils.PrintSelectedModules(data, config, contextMap)
}
