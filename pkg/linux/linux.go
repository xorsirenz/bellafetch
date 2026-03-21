package linux

import (
	"github.com/xorsirenz/bellafetch/pkg/utils"
)

func GetLinuxData() utils.Data {
	return utils.Data{
		Username:   Username(),
		Hostname:   Hostname(),
		PrettyName: PrettyName(),
		Kernel:     Kernel(),
		Uptime:     Uptime(),
		PkgManager: PkgManager(),
		WM:         "",
		Cpu:        Cpu(),
		Gpu:        Gpu(),
		DiskSpace:  Storage(),
		Memory:     Memory(),
	}
}
