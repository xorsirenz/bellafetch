package linux

import (
	"fmt"
	"strings"
	"strconv"
	"time"

	"github.com/xorsirenz/bellafetch/internal/utils"
)

func Uptime() string {
	uptimeFile := "/proc/uptime"

	data, err := utils.OpenFile(uptimeFile)
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
