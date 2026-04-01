package linux

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Kernel() string {
	versionFile := "/proc/version"

	contents, err := os.ReadFile(versionFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}

	kernelVersion := ""
	scanner := bufio.NewScanner(strings.NewReader(string(contents)))

	for scanner.Scan() {
		kernelInfo := scanner.Text()
		kernelLines := strings.Split(kernelInfo, " ")
		rgx := regexp.MustCompile(`^(\d+\.)(\d+\.)(\d+)`)
		kernelVersion = rgx.FindString(kernelLines[2])
	}
	return kernelVersion
}
