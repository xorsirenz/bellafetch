package main

import (
	"fmt"
	"runtime"
	"os"

	"github.com/xorsirenz/bellafetch/pkg/linux"
)

var data = make(map[string]any)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func checkOS() {
	goos := runtime.GOOS
	switch goos {
	case "linux":
		getLinuxData()
	case "freebsd","netbsd","openbsd", "dragonfly" :
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
} 

func getLinuxData() {
		data["user"] = linux.Username()
		data["host"] = linux.Hostname()
		data["prettyname"] = linux.PrettyName()
		data["kernel"] = linux.Kernel()
		data["uptime"] = linux.Uptime()
		data["pkgs"] = linux.PkgManager()
		data["wm"] = ""
		data["cpu"] = linux.Cpu()
		data["gpu"] = linux.Gpu()
		data["diskSpace"] = linux.Storage()
		data["memory"] = linux.Memory()
}


func main() {
	checkOS()
	clearScreen()

	fmt.Println("")
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]")
	fmt.Println("")
	fmt.Printf("  host    :: %v@%v\n", data["user"], data["host"])
	fmt.Printf("  os      :: %v\n", data["prettyname"])
	fmt.Printf("  ver     :: %v\n", data["kernel"])
	fmt.Printf("  uptime  :: %v\n", data["uptime"])
	fmt.Printf("  pkgs    :: %v\n", data["pkgs"])
	fmt.Printf("  wm      :: %v\n", data["wm"])
	fmt.Printf("  cpu     :: %v\n", data["cpu"])
	fmt.Printf("  gpu     :: %v\n", data["gpu"])
	fmt.Printf("  storage :: %v\n", data["diskSpace"])
	fmt.Printf(" memory  :: %v\n", data["memory"])
	fmt.Println("")
}
