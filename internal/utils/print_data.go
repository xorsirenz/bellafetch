package utils

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Banner() {
	const banner = `
	 bellafetch
    [github :: xorsirenz]
	`

	fmt.Println(banner)
}

func BuildSelectedModules(data interface{}, config map[string]bool) []string {
	contextMap := map[string]string{
		"Host":       "   host    ::",
		"PrettyName": "   os      ::",
		"Kernel":     "   ver     ::",
		"Uptime":     "   uptime  ::",
		"Packages":   "   pkgs    ::",
		"Shell":      "   shell   ::",
		"Terminal":   "   term    ::",
		"WM":         "   wm      ::",
		"Cpu":        "   cpu     ::",
		"Gpu":        "   gpu     ::",
		"DiskSpace":  "   storage ::",
		"Memory":     "  memory  ::",
	}

	var moduleLines []string
	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataValue.NumField(); i++ {
		moduleName := dataType.Field(i).Name

		if config[moduleName] {
			moduleValue := dataValue.Field(i).Interface()

			moduleLabel := moduleName
			if ctx, ok := contextMap[moduleName]; ok {
				moduleLabel = ctx
			}

			moduleLines = append(moduleLines, fmt.Sprintf("%s %v", moduleLabel, moduleValue))
		}
	}
	return moduleLines
}

func RenderAsciiWithSelectedModules(ascii string, moduleLines []string) {
	asciiLines := strings.Split(strings.Trim(ascii, "\n"), "\n")

	for i := range asciiLines {
		asciiLines[i] = strings.TrimRightFunc(asciiLines[i], unicode.IsSpace)
	}

	maxWidth := 0
	for _, line := range asciiLines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	asciiLen := len(asciiLines)
	moduleLen := len(moduleLines)

	offset := 0
	if asciiLen > moduleLen {
		offset = (asciiLen - moduleLen) / 2
	}

	asciiLenTotal := asciiLen
	if moduleLen > asciiLenTotal {
		asciiLenTotal = moduleLen
	}

	for i := 0; i < asciiLenTotal; i++ {
		var asciiText, moduleText string
		if i < asciiLen {
			asciiText = asciiLines[i]
		}

		textIndex := i - offset
		if textIndex >= 0 && textIndex < moduleLen {
			moduleText = moduleLines[textIndex]
		}

		padding := maxWidth - len(asciiText)
		if padding < 0 {
			padding = 0
		}

		fmt.Printf("%s%s%s\n", asciiText, strings.Repeat(" ", padding+1), moduleText)
	}
}
