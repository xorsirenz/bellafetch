package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func LoadConfig(path string) (map[string]bool, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config map[string]bool
	err = json.Unmarshal(file, &config)
	return config, err
}

func PrintSelectedFields(data interface{}, config map[string]bool, contextMap map[string]string) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name

		if config[fieldName] {
			fieldValue := val.Field(i).Interface()

			label := fieldName
			if ctx, ok := contextMap[fieldName]; ok {
				label = ctx
			}

			fmt.Printf("%s %v\n", label, fieldValue)
		}
	}
}
