package linux

import (
	"github.com/xorsirenz/bellafetch/internal/utils"
)

func GetLinuxData() utils.Data {

	osMap := OsRelease()

	return utils.Data{
		Host:   	Host(),
		PrettyName: osMap["PRETTY_NAME"],
		Kernel:     Kernel(),
		Uptime:     Uptime(),
		Packages:   PkgManager(osMap),
		Shell:		Shell(),
		Terminal:	Terminal(),
		WM:         Desktop(),
		Cpu:        Cpu(),
		Gpu:        Gpu(),
		DiskSpace:  DiskSpace(),
		Memory:     Memory(),
	}
}
