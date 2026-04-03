package linux

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var CUSTOM_PCI_IDS_PATH = ""
var CUSTOM_AMDGPU_IDS_PATH = ""

func resolvePciIDsPath() string {
	if p := os.Getenv("CUSTOM_PCI_IDS_PATH"); p != "" {
		return p
	}

	if CUSTOM_PCI_IDS_PATH != "" {
		return CUSTOM_PCI_IDS_PATH
	}

	return "/usr/share/hwdata/pci.ids"
}

func Gpu() string {
	pciDir := "/sys/bus/pci/devices"
	pciFile := resolvePciIDsPath()

	pciContents, err := os.ReadFile(pciFile)
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	idsLines := strings.Split(string(pciContents), "\n")

	devices, err := filepath.Glob(filepath.Join(pciDir, "*"))
	if err != nil {
		_ = fmt.Errorf("Error: %v", err)
	}
	for _, devicePath := range devices {
		vendorID, err := os.ReadFile(filepath.Join(devicePath, "vendor"))
		if err != nil {
			_ = fmt.Errorf("Error: %v", err)
		}
		deviceID, err := os.ReadFile(filepath.Join(devicePath, "device"))
		if err != nil {
			_ = fmt.Errorf("Error: %v", err)
		}

		vendorStr := strings.TrimPrefix(strings.TrimSpace(string(vendorID)), "0x")
		deviceStr := strings.TrimPrefix(strings.TrimSpace(string(deviceID)), "0x")

		classFile := filepath.Join(devicePath, "class")
		classData, err := os.ReadFile(classFile)
		if err != nil {
			_ = fmt.Errorf("Error: %v", err)
		}
		class := strings.TrimSpace(string(classData))
		if !strings.HasPrefix(class, "0x0300") {
			continue
		}

		var currentVendor, vendorName, deviceName string

		for _, line := range idsLines {
			if strings.HasPrefix(line, vendorStr+" ") {
				parts := strings.SplitN(line, " ", 2)
				if len(parts) == 2 {
					vendorName = strings.TrimSpace(parts[1])
					currentVendor = vendorStr
				}
			} else if strings.HasPrefix(line, "\t") && currentVendor == vendorStr {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, deviceStr+" ") {
					parts := strings.SplitN(line, " ", 2)
					if len(parts) == 2 {
						deviceName = strings.TrimSpace(parts[1])
					}
				}
			}
		}
		if vendorName != "" && deviceName != "" {
			return fmt.Sprintf("%s %s", vendorName, deviceName)
		}
	}
	return ""
}
