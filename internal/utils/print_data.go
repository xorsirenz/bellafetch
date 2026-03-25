package utils

import (
	"fmt"
	"reflect"
)

func PrintSelectedModules(data interface{}, config map[string]bool, contextMap map[string]string) {
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
