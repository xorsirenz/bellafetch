package linux

import (
	"fmt"
	"syscall"

	"github.com/xorsirenz/bellafetch/internal/utils"
)
func DiskSpace() string {
	path := "/"
	var fs syscall.Statfs_t
	err := syscall.Statfs(path, &fs)
	if err != nil {
		fmt.Println("Error:", err)
	}
	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := total - free

	totalConverted := utils.PrettyByteSize(total)
	usedConverted := utils.PrettyByteSize(used)

	return fmt.Sprintf("%s / %s", usedConverted, totalConverted)
}
