package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Ascii []Ascii `json:"Ascii"`
	Modules []Modules `json:"Modules"`
}

type Ascii struct {
	Ascii bool `json:"Ascii"`
}

type Modules struct {
	Host       bool `json:"Host"`
	PrettyName bool `json:"PrettyName"`
	Kernel     bool `json:"Kernel"`
	Uptime     bool `json:"Uptime"`
	Package    bool `json:"Packages"`
	Shell      bool `json:"Shell"`
	Terminal   bool `json:"Terminal"`
	WM         bool `json:"WM"`
	Cpu        bool `json:"Cpu"`
	Gpu        bool `json:"Gpu"`
	DiskSpace  bool `json:"DiskSpace"`
	Memory     bool `json:"Memory"`
}

func LoadConfig() (map[string]bool, map[string]bool) {
	configFile, err := configDirExists()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	var configData struct {
		Ascii []map[string]bool `json:"Ascii"`
		Modules []map[string]bool `json:"Modules"`
	}

	err = json.Unmarshal(file, &configData)
	if err != nil {
		log.Fatalf("Cannot unmarshal json: %v", err)
	}

	return configData.Modules[0], configData.Ascii[0]
}

func configDirExists() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", fmt.Errorf("Error failed to find user config direcoty: %v", err)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		createDefaultConfig(configPath)
	}
	return configPath, err
}

func getConfigPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find user config directory: %w", err)
	}
	return filepath.Join(userConfigDir, "bellafetch", "config"), err
}

func createDefaultConfig(configPath string) error {
	configDir := filepath.Dir(configPath)

	configJson := Config{
		Ascii : []Ascii{
			{ Ascii: true },
		},

		Modules: []Modules{

			{
				Host:       true,
				PrettyName: true,
				Kernel:     true,
				Uptime:     true,
				Package:    true,
				Shell:      true,
				Terminal:   true,
				WM:         true,
				Cpu:        true,
				Gpu:        true,
				DiskSpace:  true,
				Memory:     true,
			},
		},
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("Failed to create config directory: %w", err)
	}

	defaultConfigData, err := json.MarshalIndent(configJson, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to Marshal config file: %w", err)
	}
	if err := os.WriteFile(configPath, defaultConfigData, 0644); err != nil {
		return fmt.Errorf("Failed to write default config file: %w", err)
	}
	return err
}
