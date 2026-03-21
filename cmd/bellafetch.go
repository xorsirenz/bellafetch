package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xorsirenz/bellafetch/pkg/linux"
	"github.com/xorsirenz/bellafetch/pkg/utils"
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
		fmt.Println("Error: Bellafetch is not capitable with any BSD derivatives right now..")
		os.Exit(-1)
	case "darwin":
		fmt.Println("Error: Bellafetch is not capitable with Darwin/Mac OSX right now..")
		os.Exit(-1)
	case "windows":
		fmt.Println("Error: Bellafetch is not capitable with Windows right now..")
		os.Exit(-1)
	default:
		fmt.Println("Error: Bellafetch cannot detect OS target..")
		os.Exit(-1)
	}
	return utils.Data{}
}

func main() {
	data := checkOS()
	clearScreen()

	fmt.Println("")
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]")
	fmt.Println("")
	fmt.Printf("  host    :: %v@%v\n", data.Username, data.Hostname)
	fmt.Printf("  os      :: %v\n", data.PrettyName)
	fmt.Printf("  ver     :: %v\n", data.Kernel)
	fmt.Printf("  uptime  :: %v\n", data.Uptime)
	fmt.Printf("  pkgs    :: %v\n", data.PkgManager)
	fmt.Printf("  wm      :: %v\n", data.WM)
	fmt.Printf("  cpu     :: %v\n", data.Cpu)
	fmt.Printf("  gpu     :: %v\n", data.Gpu)
	fmt.Printf("  storage :: %v\n", data.DiskSpace)
	fmt.Printf(" memory  :: %v\n", data.Memory)
	fmt.Println("")
}
