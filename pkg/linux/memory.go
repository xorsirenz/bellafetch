package linux

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xorsirenz/bellafetch/internal/utils"
)

func Memory() string {
	meminfoFile := "/proc/meminfo"

	contents, err := os.ReadFile(meminfoFile)
	if err != nil {
		_ = fmt.Errorf("Error:", err)
	}

	var memTotal, memAvailable uint64
	scanner := bufio.NewScanner(strings.NewReader(string(contents)))
	for scanner.Scan() {
		memInfo := scanner.Text()
		if strings.HasPrefix(memInfo, "MemTotal:") {
			fields := strings.Fields(memInfo)
			memValue, _ := strconv.ParseUint(fields[1], 10, 64)
			memTotal = memValue * 1024
		} else if strings.HasPrefix(memInfo, "MemAvailable:") {
			fields := strings.Fields(memInfo)
			memValue, _ := strconv.ParseUint(fields[1], 10, 64)
			memAvailable = memValue * 1024
		}
	}

	memUsed := memTotal - memAvailable
	memTotalPretty := utils.PrettyByteSize(memTotal)
	memUsedPretty := utils.PrettyByteSize(memUsed)
	return fmt.Sprintf("%s / %s", memUsedPretty, memTotalPretty)
}
