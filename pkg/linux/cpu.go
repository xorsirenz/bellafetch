package linux

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

func Cpu() string {
	cpuinfoFile := "/proc/cpuinfo"

	contents, err := os.ReadFile(cpuinfoFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var cpuVersion string
	scanner := bufio.NewScanner(strings.NewReader(string(contents)))

	for scanner.Scan() {
		cpuInfo := scanner.Text()
		if strings.HasPrefix(cpuInfo, "model name") {
			cpu := strings.SplitN(cpuInfo, ":", 2)
			if len(cpu) == 2 {
				cpuVersion = strings.TrimSpace(cpu[1])
			}
		}
	}
	return cpuVersion
}
