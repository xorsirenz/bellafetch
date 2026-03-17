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
		fmt.Println("bellafetch is only capitable with Linux right now..")
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
		os.Exit(-1)
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

func vga() string {
    pciDir := "/sys/bus/pci/devices"
    idsFile := "/usr/share/hwdata/pci.ids"

    idsContents, err := openFile(idsFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading pci.ids: %v\n", err)
		os.Exit(-1)
    }
    idsLines := strings.Split(string(idsContents), "\n")

    devices, err := filepath.Glob(filepath.Join(pciDir, "*"))
	if err != nil {
		fmt.Println("Glob error:", err)
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

func username() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-1)
	}
	return currentUser.Username
}

func hostname() string {
	hostnameFile := "/etc/hostname"

	hostnameContents, err := openFile(hostnameFile)
	if err != nil {
		fmt.Println("error:", err)
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
		os.Exit(-1)
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
		fmt.Println("apt not supported yet")
	default:
		fmt.Println("No supported package manager detected")
	}
	return 0
}

func pacman() int {
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
	uptimeFile := "/proc/uptime"

    data, err := openFile(uptimeFile)
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

func memory() string {
	meminfoFile := "/proc/meminfo"

    contents, err:= openFile(meminfoFile)
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

func main(){
	checkOS()
	clearScreen()
	fmt.Println("")
	fmt.Println("	bellafetch")
	fmt.Println("  [github : xorsirenz]")
	fmt.Println("")
	fmt.Println("  host    ::", username() + "@" + hostname())
	fmt.Println("  os      ::", prettyName())
	fmt.Println("  ver     ::", kernel())
	fmt.Println("  uptime  ::", uptime()) 
	fmt.Println("  pkgs    ::", pkgManager())
	fmt.Println("  wm      ::",) 
	fmt.Println("  cpu     ::", cpu()) 
	fmt.Println("  gpu     ::", vga()) 
	fmt.Println("  storage ::",) 
	fmt.Println(" memory  ::", memory()) 
	fmt.Println("")
}
