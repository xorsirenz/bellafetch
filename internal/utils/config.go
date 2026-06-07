package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Banner      bool            `yaml:"Banner"`
	Ascii       string          `yaml:"Ascii"`
	Modules     map[string]bool `yaml:"Modules"`
	ColorBlocks bool            `yaml:"ColorBlocks"`
}

func LoadConfig() Config {
	configFile, err := configDirExists()
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	file, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Cannot unmarshal yaml: %v", err)
	}

	return config
}

func configDirExists() (string, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return "", fmt.Errorf("Failed to find user config direcoty: %v", err)
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := createDefaultConfig(configPath); err != nil {
			return "", err
		}
	}
	return configPath, nil
}

func getConfigPath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("Failed to find user config directory: %w", err)
	}
	return filepath.Join(userConfigDir, "bellafetch", "config.yaml"), nil
}

func createDefaultConfig(configPath string) error {
	configDir := filepath.Dir(configPath)

	configYaml := Config{
		Banner: true,
		Ascii:  "none",
		Modules: map[string]bool{
			"Host":       true,
			"PrettyName": true,
			"Kernel":     true,
			"Uptime":     true,
			"Packages":   true,
			"Shell":      true,
			"Terminal":   true,
			"WM":         true,
			"Cpu":        true,
			"Gpu":        true,
			"DiskSpace":  true,
			"Memory":     true,
		},
		ColorBlocks: true,
	}

	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("Failed to create config directory: %w", err)
	}

	defaultConfigData, err := yaml.Marshal(&configYaml)
	if err != nil {
		return fmt.Errorf("Failed to marshal config file: %w", err)
	}

	if err := os.WriteFile(configPath, defaultConfigData, 0644); err != nil {
		return fmt.Errorf("Failed to write default config file: %w", err)
	}
	return nil
}
