package linux

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
)
func Memory() string {
	meminfoFile := "/proc/meminfo"

	contents, err := OpenFile(meminfoFile)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var memTotal, memAvailable uint64
	scanner := bufio.NewScanner(strings.NewReader(contents))
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
	memTotalPretty := prettyByteSize(memTotal)
	memUsedPretty := prettyByteSize(memUsed)
	return fmt.Sprintf("%s / %s", memUsedPretty, memTotalPretty)
}
