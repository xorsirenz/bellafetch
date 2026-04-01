package linux

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Uptime() string {
	uptimeFile := "/proc/uptime"

	data, err := os.ReadFile(uptimeFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		_ = fmt.Errorf("Error: %v", err)
	}

	seconds, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	realtime := time.Duration(seconds) * time.Second

	days := int(realtime.Hours() / 24)
	hours := int(realtime.Hours()) % 24
	minutes := int(realtime.Minutes()) % 60
	return fmt.Sprintf("%dd %02dh %02dm", days, hours, minutes)

}
