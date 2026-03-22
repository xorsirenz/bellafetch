package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type Config struct {
	Host       bool `json:"Host"`
	PrettyName bool `json:"PrettyName"`
	Kernel     bool `json:"Kernel"`
	Uptime     bool `json:"Uptime"`
	Package    bool `json:"Packages"`
	WM         bool `json:"WM"`
	Cpu        bool `json:"Cpu"`
	Gpu        bool `json:"Gpu"`
	DiskSpace  bool `json:"DiskSpace"`
	Memory     bool `json:"Memory"`
}

func LoadConfig() (map[string]bool, error) {
	configFile, err := configDirExists()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(1)
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var config map[string]bool
	err = json.Unmarshal(file, &config)
	return config, err
}

func configDirExists() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find user config directory: %w", err)
	}
	bellafetchConfigPath := filepath.Join(configDir, "bellafetch")

	if err := os.MkdirAll(bellafetchConfigPath, 0755); err != nil {
		return "", fmt.Errorf("Failed to create config directory: %w", err)
	}

	configJson := Config{
		Host:       true,
		PrettyName: true,
		Kernel:     true,
		Uptime:     true,
		Package:    true,
		WM:         false,
		Cpu:        true,
		Gpu:        true,
		DiskSpace:  true,
		Memory:     true,
	}

	configFilePath := filepath.Join(bellafetchConfigPath, "config")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		defaultConfigData, err := json.MarshalIndent(configJson, "", "  ")
		if err != nil {
			return "", fmt.Errorf("Failed to Marshal config file: %w", err)
		}
		if err := os.WriteFile(configFilePath, defaultConfigData, 0644); err != nil {
			return "", fmt.Errorf("Failed to write default config file: %w", err)
		}
	}
	return configFilePath, err
}

func PrintSelectedFields(data interface{}, config map[string]bool, contextMap map[string]string) {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	for i := range val.NumField() {
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
