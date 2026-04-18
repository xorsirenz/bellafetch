package utils

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func PrintData(ascii string, data interface{}, config Config) {
	clearScreen()
	printBanner()
	selectedModules := buildSelectedModules(data, config)
	renderAsciiWithSelectedModules(ascii, selectedModules)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printBanner() {
	const banner = `
	 bellafetch
    [github :: xorsirenz]`
	fmt.Println(banner)
}

func buildSelectedModules(data interface{}, config Config) []string {
	contextMap := getContextMap()
	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)

	var moduleLines []string
	for i := 0; i < dataValue.NumField(); i++ {
		moduleName := dataType.Field(i).Name

		if !config.Modules[moduleName] {
			continue
		}

		moduleLabel := getModuleLabel(moduleName, contextMap)
		moduleValue := dataValue.Field(i).Interface()

		moduleLines = append(moduleLines, formatModule(moduleLabel, moduleValue))
	}
	return moduleLines
}

func getContextMap() map[string]string {
	return map[string]string{
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
}

func getModuleLabel(moduleName string, contextMap map[string]string) string {
	if ctx, ok := contextMap[moduleName]; ok {
		return ctx
	}
	return moduleName
}

func formatModule(label string, value interface{}) string {
	return fmt.Sprintf("%s %v", label, value)
}

func renderAsciiWithSelectedModules(ascii string, moduleLines []string) {
	asciiLines := prepareAsciiLines(ascii)
	maxWidth := getMaxWidth(asciiLines)

	asciiLen := len(asciiLines)
	moduleLen := len(moduleLines)
	offset := calculateOffset(asciiLen, moduleLen)

	for i := 0; i < max(asciiLen, moduleLen); i++ {
		asciiText := getAsciiText(i, asciiLines)
		moduleText := getModuleText(i, moduleLines, offset)
		printLine(asciiText, moduleText, maxWidth)
	}
}

func prepareAsciiLines(ascii string) []string {
	asciiLines := strings.Split(strings.Trim(ascii, "\n"), "\n")
	for i := range asciiLines {
		asciiLines[i] = strings.TrimRightFunc(asciiLines[i], unicode.IsSpace)
	}
	return asciiLines
}

func getMaxWidth(asciiLines []string) int {
	maxWidth := 0
	for _, line := range asciiLines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	return maxWidth
}

func calculateOffset(asciiLen, moduleLen int) int {
	if asciiLen > moduleLen {
		return (asciiLen - moduleLen) / 2
	}
	return 0
}

func getAsciiText(index int, asciiLines []string) string {
	if index < len(asciiLines) {
		return asciiLines[index]
	}
	return ""
}

func getModuleText(index int, moduleLines []string, offset int) string {
	textIndex := index - offset
	if textIndex >= 0 && textIndex < len(moduleLines) {
		return moduleLines[textIndex]
	}
	return ""
}

func printLine(asciiText, moduleText string, maxWidth int) {
	padding := maxWidth - len(asciiText)
	if padding < 0 {
		padding = 0
	}
	fmt.Printf("%s%s%s\n", asciiText, strings.Repeat(" ", padding+1), moduleText)
}
