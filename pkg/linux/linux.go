package linux

import (
	"github.com/xorsirenz/bellafetch/internal/utils"
)

func GetLinuxData() utils.Data {

	osMap := OsRelease()
	id := GetID(osMap)

	return utils.Data{
		Id:         id,
		IdLike:		osMap["ID_LIKE"],
		Host:       Host(),
		PrettyName: osMap["PRETTY_NAME"],
		Kernel:     Kernel(),
		Uptime:     Uptime(),
		Packages:   PkgManager(GetIDLike(osMap)),
		Shell:      Shell(),
		Terminal:   Terminal(),
		WM:         Desktop(),
		Cpu:        Cpu(),
		Gpu:        Gpu(),
		DiskSpace:  DiskSpace(),
		Memory:     Memory(),
	}
}
