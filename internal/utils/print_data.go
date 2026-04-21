package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

func PrintData(ascii string, data interface{}, config Config) {
	banner := []string{
		"         bellafetch",
		"    [github :: xorsirenz]",
	}

	clearScreen()
	selectedModules := buildSelectedModules(data, config)
	renderLayout(ascii, selectedModules, banner)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func buildSelectedModules(data interface{}, config Config) []string {
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

	dataValue := reflect.ValueOf(data)
	dataType := reflect.TypeOf(data)

	var moduleLines []string
	for i := 0; i < dataValue.NumField(); i++ {
		moduleName := dataType.Field(i).Name

		if !config.Modules[moduleName] {
			continue
		}

		moduleValue := dataValue.Field(i).Interface()
		moduleLabel := moduleName
		if ctx, ok := contextMap[moduleName]; ok {
			moduleLabel = ctx
		}

		moduleLines = append(moduleLines,
			fmt.Sprintf("%s %v", moduleLabel, moduleValue))
	}
	return moduleLines
}

func renderLayout(ascii string, moduleLines []string, banner []string) {
	asciiLines := prepareAscii(ascii)
	maxWidth := getMaxWidth(asciiLines)

	rightColumn := append(banner, moduleLines...)

	asciiLen := len(asciiLines)
	columnLen := len(rightColumn)

	maxHeight := max(asciiLen, columnLen)

	asciiOffset := (maxHeight - asciiLen) / 2
	rightOffset := (maxHeight - columnLen) / 2

	for i := 0; i < maxHeight; i++ {
		asciiText := ""
		if i >= asciiOffset && i < asciiOffset+asciiLen {
			asciiText = asciiLines[i-asciiOffset]
		}

		columnText := ""
		if i >= rightOffset && i < rightOffset+columnLen {
			columnText = rightColumn[i-rightOffset]
		}
		printLine(asciiText, columnText, maxWidth)
	}
}

func prepareAscii(ascii string) []string {
	asciiLines := strings.Split(strings.Trim(ascii, "\n"), "\n")
	for i := range asciiLines {
		asciiLines[i] = strings.TrimRightFunc(asciiLines[i], unicode.IsSpace)
	}
	return asciiLines
}

func getMaxWidth(asciiLines []string) int {
	maxWidth := 0

	for _, line := range asciiLines {
		asciiWidth := displayWidth(line)
		if asciiWidth > maxWidth {
			maxWidth = asciiWidth
		}
	}
	return maxWidth
}

func stripANSI(asciiString string) string {
	var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

	return ansiRegex.ReplaceAllString(asciiString, "")
} 

func displayWidth(asciiString string) int {
	asciiStr := stripANSI(asciiString)
	width := 0

	for _, r := range asciiStr {
		if r <= 127 {
			width++
			continue
		}
		// Common emoji ranges
		if (r >= 0x1F300 && r <= 0x1FAFF) ||
			(r >= 0x1F1E6 && r <= 0x1F1FF) {
			width += 2
			continue
		}
		// CJK / fullwidth chars
		if (r >= 0x1100 && r <= 0x115F) ||
			(r >= 0x2329 && r <= 0x232A) ||
			(r >= 0x2E80 && r <= 0xA4CF) ||
			(r >= 0xAC00 && r <= 0xD7A3) ||
			(r >= 0xF900 && r <= 0xFAFF) {
			width += 2
			continue
		}
		if (r >= 0x2500 && r <= 0x257F) || // box drawing
			(r >= 0x2580 && r <= 0x259F) { // block elements
			width += 1
			continue
		}
		width++
	}
	return width
}

func padRight(asciiString string, targetWidth int) string {
	currentWidth := displayWidth(asciiString)
	if currentWidth >= targetWidth {
		return asciiString
	}
	return asciiString + strings.Repeat(" ", targetWidth-currentWidth)
}

func printLine(asciiText, rightText string, maxWidth int) {
	const gap = 2

	fmt.Printf("%s%s%s\n", 
		padRight(asciiText, maxWidth),
		strings.Repeat(" ", gap),
		rightText,
	)
}
