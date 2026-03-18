package linux

import (
	"fmt"
	"strings"
	"regexp"
	"bufio"
)

func Kernel() string {
	versionFile := "/proc/version"

	contents, err := OpenFile(versionFile)
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
