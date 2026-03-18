package linux

import (
	"fmt"
	"strings"
	"path/filepath"
)

func Vga() string {
	pciDir := "/sys/bus/pci/devices"
	idsFile := "/usr/share/hwdata/pci.ids"

	idsContents, err := OpenFile(idsFile)
	if err != nil {
		fmt.Println("Error:", err)
	}
	idsLines := strings.Split(string(idsContents), "\n")

	devices, err := filepath.Glob(filepath.Join(pciDir, "*"))
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, devPath := range devices {
		vendorID, err := OpenFile(filepath.Join(devPath, "vendor"))
		if err != nil {
			continue
		}
		deviceID, err := OpenFile(filepath.Join(devPath, "device"))
		if err != nil {
			continue
		}

		vendorStr := strings.TrimPrefix(strings.TrimSpace(string(vendorID)), "0x")
		deviceStr := strings.TrimPrefix(strings.TrimSpace(string(deviceID)), "0x")

		classFile := filepath.Join(devPath, "class")
		classData, err := OpenFile(classFile)
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

