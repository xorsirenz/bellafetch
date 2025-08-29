package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"regexp"
	"strings"
)

func username() string {
	current_user, err := user.Current()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	return current_user.Username
}

func hostname() string {
	hostname_file, err := os.ReadFile("/etc/hostname")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}

	hostname := string(hostname_file)
	host := strings.TrimSuffix(hostname, "\n")
	return host
}

func distro() string {
	os_file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	defer os_file.Close()
	
	pretty_name := ""
	scanner := bufio.NewScanner(os_file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "PRETTY_NAME") {
			pretty_info := scanner.Text()
			pretty_name = pretty_info[13 : len(pretty_info)-1]
		}
	}
	return pretty_name
}

func kernel() string {
	kernel_file, err := os.Open("/proc/version")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer kernel_file.Close()

	kernel_version := ""
	scanner := bufio.NewScanner(kernel_file)

	for scanner.Scan() {
		kernel_info := scanner.Text()
		kernel_lines := strings.Split(kernel_info, " ")
		rgx := regexp.MustCompile(`^(\d+\.)?(\d+\.)?(\*|\d+)`)
		kernel_version = rgx.FindString(kernel_lines[2])
	}
	return kernel_version
}

func main(){
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]\n")
	fmt.Println("  host    ::", username() + "@" + hostname())
	fmt.Println("  os	   ::", distro())
	fmt.Println("  ver	   ::", kernel())
	fmt.Println("  uptime  ::",) 
	fmt.Println("  pkgs    ::",)
	fmt.Println("  wm      ::",) 
	fmt.Println("  cpu     ::",) 
	fmt.Println("  gpu     ::",) 
	fmt.Println("  storage ::",) 
	fmt.Println("  mem    ::",) 
}
