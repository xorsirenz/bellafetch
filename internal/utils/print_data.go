package utils

import (
	"fmt"
	"reflect"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Banner() {
	const banner = `
	 bellafetch
    [github : xorsirenz]
	`

	fmt.Println(banner)
}

func PrintSelectedModules(data interface{}, config map[string]bool) {
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

	for i := range dataValue.NumField() {
		moduleName := dataType.Field(i).Name

		if config[moduleName] {
			moduleValue := dataValue.Field(i).Interface()

			ModuleLabel := moduleName
			if ctx, ok := contextMap[moduleName]; ok {
				ModuleLabel = ctx
			}

			fmt.Printf("%s %v\n", ModuleLabel, moduleValue)
		}
	}
}
