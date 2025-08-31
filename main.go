package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"
)

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
			prettyName = prettyInfo[13 : len(prettyInfo)-1]
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

func main(){
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]\n")
	fmt.Println("  host    ::", username() + "@" + hostname())
	fmt.Println("  os	   ::", distro())
	fmt.Println("  ver	   ::", kernel())
	fmt.Println("  uptime  ::",) 
	fmt.Println("  pkgs    ::", packages())
	fmt.Println("  wm      ::",) 
	fmt.Println("  cpu     ::",) 
	fmt.Println("  gpu     ::",) 
	fmt.Println("  storage ::",) 
	fmt.Println("  mem    ::",) 
}
