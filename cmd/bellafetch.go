package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xorsirenz/bellafetch/pkg/linux"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("Error: Bellafetch is only capitable with Linux right now..")
		os.Exit(-1)
	}
}

func main() {
	checkOS()
	clearScreen()

	user := linux.Username()
	host := linux.Hostname()
	prettyName := linux.PrettyName()
	kernel := linux.Kernel()
	uptime := linux.Uptime()
	pkgs := linux.PkgManager()
	wm := ""
	cpu := linux.Cpu()
	vga := linux.Vga()
	storage := linux.Storage()
	memory := linux.Memory()

	fmt.Println("")
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]")
	fmt.Println("")
	fmt.Println("  host    ::", user+"@"+host)
	fmt.Println("  os      ::", prettyName)
	fmt.Println("  ver     ::", kernel)
	fmt.Println("  uptime  ::", uptime)
	fmt.Println("  pkgs    ::", pkgs)
	fmt.Println("  wm      ::", wm)
	fmt.Println("  cpu     ::", cpu)
	fmt.Println("  gpu     ::", vga)
	fmt.Println("  storage ::", storage)
	fmt.Println(" memory  ::", memory)
	fmt.Println("")
}
