package linux

import (
	"github.com/xorsirenz/bellafetch/internal/utils"
)

func GetLinuxData() utils.Data {
	return utils.Data{
		Host:   	Host(),
		PrettyName: PrettyName(),
		Kernel:     Kernel(),
		Uptime:     Uptime(),
		Packages:   PkgManager(),
		Shell:		Shell(),
		Terminal:	Terminal(),
		WM:         Wm(),
		Cpu:        Cpu(),
		Gpu:        Gpu(),
		DiskSpace:  DiskSpace(),
		Memory:     Memory(),
	}
}
