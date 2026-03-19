package linux

import (
	"fmt"
	"bufio"
	"strings"

	"github.com/xorsirenz/bellafetch/pkg/utils"
)
func Cpu() string {
	cpuinfoFile := "/proc/cpuinfo"

	contents, err := utils.OpenFile(cpuinfoFile)
	if err != nil {
		fmt.Println("Error:", err)
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
