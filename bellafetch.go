package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func clearScreen() {
	clear, _ := exec.Command("clear").Output()
	os.Stdout.Write(clear)
}

func checkOS() {
	if runtime.GOOS != "linux" {
		fmt.Println("bellafetch is only capitable with Linux right now..")
		os.Exit(-1)
	}
}

func username() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	return currentUser.Username
}

func hostname() string {
	hostnameFile, err := os.ReadFile("/etc/hostname")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}

	hostname := string(hostnameFile)
	host := strings.TrimSuffix(hostname, "\n")
	return host
}

func distro() string {
	osFile, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	defer osFile.Close()
	
	prettyName := ""
	scanner := bufio.NewScanner(osFile)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME") {
			prettyInfo := scanner.Text()
			kv := strings.Split(prettyInfo, "=")
			prettyName = strings.Trim(kv[1], "\"")
		}
	}
	return prettyName
}

func kernel() string {
	kernelFile, err := os.Open("/proc/version")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer kernelFile.Close()

	kernelVersion := ""
	scanner := bufio.NewScanner(kernelFile)

	for scanner.Scan() {
		kernelInfo := scanner.Text()
		kernelLines := strings.Split(kernelInfo, " ")
		rgx := regexp.MustCompile(`^(\d+\.)(\d+\.)(\d+)`)
		kernelVersion = rgx.FindString(kernelLines[2])
	}
	return kernelVersion
}

func packages() int {
	out, err:= exec.Command("pacman", "-Q").Output()
	if err != nil {
		fmt.Println(err)
	}

	output := string(out)
	outputLines := strings.Split(output, "\n")
	lines := len(outputLines) -1
	return lines
}

func uptime() string {
    data, err := os.ReadFile("/proc/uptime")
    if err != nil {
		fmt.Println(err)
    }

    fields := strings.Fields(string(data))
    if len(fields) < 1 {
		fmt.Println(err)
    }

    seconds, err := strconv.ParseFloat(fields[0], 64)
    if err != nil {
		fmt.Println(err)
    }

	realtime := time.Duration(seconds) * time.Second

	days := int(realtime.Hours() / 24)
	hours := int(realtime.Hours()) % 24
	minutes := int(realtime.Minutes()) % 60
	return fmt.Sprintf("%dd %02dh %02dm", days, hours, minutes)

}

func cpu() string {
	cpuinfoFile, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer cpuinfoFile.Close()

	cpuVersion := ""
	scanner := bufio.NewScanner(cpuinfoFile)

	for scanner.Scan() {
		cpuInfo := scanner.Text()
		if strings.HasPrefix(cpuInfo, "model name") {
			cpu := strings.SplitN(cpuInfo, ":", 2)
			if len(cpu) == 2 {
				cpuVersion := strings.TrimSpace(cpu[1])
				return cpuVersion
			}
		}
	}
	return cpuVersion
}

func main(){
	checkOS()
	clearScreen()
	fmt.Println("")
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]")
	fmt.Println("")
	fmt.Println("  host    ::", username() + "@" + hostname())
	fmt.Println("  os	   ::", distro())
	fmt.Println("  ver	   ::", kernel())
	fmt.Println("  uptime  ::",uptime()) 
	fmt.Println("  pkgs    ::", packages())
	fmt.Println("  wm      ::",) 
	fmt.Println("  cpu     ::",cpu()) 
	fmt.Println("  gpu     ::",) 
	fmt.Println("  storage ::",) 
	fmt.Println("  mem    ::",) 
}
