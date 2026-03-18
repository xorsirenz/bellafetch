package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
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
		fmt.Println("Error: Bellafetch is only capitable with Linux right now..")
		os.Exit(-1)
	}
}

func openFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func osRelease() map[string]string {
	osReleaseFile := "/etc/os-release"

	contents, err := openFile(osReleaseFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	entries := strings.Split(contents, "\n")
	osMap := make(map[string]string)

	for _, entry := range entries {
		parts := strings.Split(entry, "=")
		if len(parts) == 2 {
			osMap[parts[0]] = parts[1]
		}
	}
	return osMap
}

func username() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error:", err)
	}
	return currentUser.Username
}

func hostname() string {
	hostnameFile := "/etc/hostname"

	hostnameContents, err := openFile(hostnameFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	hostname := string(hostnameContents)
	host := strings.TrimSuffix(hostname, "\n")
	return host
}

func prettyName() string {
	prettyName := osRelease()
	return strings.Trim(prettyName["PRETTY_NAME"], "\"")
}

func kernel() string {
	versionFile := "/proc/version"

	contents, err := openFile(versionFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	kernelVersion := ""
	scanner := bufio.NewScanner(strings.NewReader(contents))

	for scanner.Scan() {
		kernelInfo := scanner.Text()
		kernelLines := strings.Split(kernelInfo, " ")
		rgx := regexp.MustCompile(`^(\d+\.)(\d+\.)(\d+)`)
		kernelVersion = rgx.FindString(kernelLines[2])
	}
	return kernelVersion
}

func pkgManager() int {
	id := osRelease()
	detectPackageMgr := id["ID"]

	switch detectPackageMgr {
	case "arch", "manjaro":
		pkgs := pacman()
		return pkgs
	case "debian", "linuxmint", "ubuntu":
		pkgs := apt()
		return pkgs
	default:
		fmt.Println("No supported package manager detected")
	}
	return 0
}

func apt() int {
	out, err := exec.Command("dpkg-query", "--list").Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	output := string(out)
	outputLines := strings.Split(output, "\n")
	count := 0
	for _, line := range outputLines {
		if strings.HasPrefix(strings.TrimSpace(line), "ii") {
			count++
		}
	}
	return count
}

func pacman() int {
	out, err := exec.Command("pacman", "-Q").Output()
	if err != nil {
		fmt.Println("Error:", err)
	}

	output := string(out)
	outputLines := strings.Split(output, "\n")
	lines := len(outputLines) - 1
	return lines
}

func uptime() string {
	uptimeFile := "/proc/uptime"

	data, err := openFile(uptimeFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		fmt.Println("Error:", err)
	}

	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	realtime := time.Duration(seconds) * time.Second

	days := int(realtime.Hours() / 24)
	hours := int(realtime.Hours()) % 24
	minutes := int(realtime.Minutes()) % 60
	return fmt.Sprintf("%dd %02dh %02dm", days, hours, minutes)

}

func cpu() string {
	cpuinfoFile := "/proc/cpuinfo"

	contents, err := openFile(cpuinfoFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	cpuVersion := ""
	scanner := bufio.NewScanner(strings.NewReader(contents))

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

func vga() string {
	pciDir := "/sys/bus/pci/devices"
	idsFile := "/usr/share/hwdata/pci.ids"

	idsContents, err := openFile(idsFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	idsLines := strings.Split(string(idsContents), "\n")

	devices, err := filepath.Glob(filepath.Join(pciDir, "*"))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	for _, devPath := range devices {
		vendorID, err := openFile(filepath.Join(devPath, "vendor"))
		if err != nil {
			continue
		}
		deviceID, err := openFile(filepath.Join(devPath, "device"))
		if err != nil {
			continue
		}

		vendorStr := strings.TrimPrefix(strings.TrimSpace(string(vendorID)), "0x")
		deviceStr := strings.TrimPrefix(strings.TrimSpace(string(deviceID)), "0x")

		classFile := filepath.Join(devPath, "class")
		classData, err := openFile(classFile)
		class := strings.TrimSpace(string(classData))
		if !strings.HasPrefix(class, "0x0300") {
			continue
		}

		var currentVendor, vendorName, deviceName string

		for _, line := range idsLines {
			if strings.HasPrefix(line, vendorStr+" ") {
				parts := strings.SplitN(line, " ", 2)
				if len(parts) == 2 {
					vendorName = strings.TrimSpace(parts[1])
					currentVendor = vendorStr
				}
			} else if strings.HasPrefix(line, "\t") && currentVendor == vendorStr {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, deviceStr+" ") {
					parts := strings.SplitN(line, " ", 2)
					if len(parts) == 2 {
						deviceName = strings.TrimSpace(parts[1])
					}
				}
			}
		}
		if vendorName != "" && deviceName != "" {
			return fmt.Sprintf("%s %s", vendorName, deviceName)
		}
	}
	return ""
}

func memory() string {
	meminfoFile := "/proc/meminfo"

	contents, err := openFile(meminfoFile)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	var memTotal, memAvailable uint64
	scanner := bufio.NewScanner(strings.NewReader(contents))
	for scanner.Scan() {
		memInfo := scanner.Text()
		if strings.HasPrefix(memInfo, "MemTotal:") {
			fields := strings.Fields(memInfo)
			memValue, _ := strconv.ParseUint(fields[1], 10, 64)
			memTotal = memValue / 1024
		} else if strings.HasPrefix(memInfo, "MemAvailable:") {
			fields := strings.Fields(memInfo)
			memValue, _ := strconv.ParseUint(fields[1], 10, 64)
			memAvailable = memValue / 1024
		}
	}

	memUsed := memTotal - memAvailable
	return fmt.Sprintf("%dMib / %dMib", memUsed, memTotal)
}

func main() {
	checkOS()
	clearScreen()

	user := username()
	host := hostname()
	prettyName := prettyName()
	kernel := kernel()
	uptime := uptime()
	pkgs := pkgManager()
	wm := ""
	cpu := cpu()
	vga := vga()
	storage := ""
	memory := memory()

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
