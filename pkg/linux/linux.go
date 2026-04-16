package linux

import (
	"github.com/xorsirenz/bellafetch/internal/utils"
)

func GetLinuxData() utils.Data {

	osMap := OsRelease()
	id := GetID(osMap)

	return utils.Data{
		Id:         id,
		Host:       Host(),
		PrettyName: osMap["PRETTY_NAME"],
		Kernel:     Kernel(),
		Uptime:     Uptime(),
		Packages:   PkgManager(id),
		Shell:      Shell(),
		Terminal:   Terminal(),
		WM:         Desktop(),
		Cpu:        Cpu(),
		Gpu:        Gpu(),
		DiskSpace:  DiskSpace(),
		Memory:     Memory(),
	}
}
