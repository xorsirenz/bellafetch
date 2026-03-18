package linux

import (
	"fmt"
	"syscall"
)
func Storage() string {
	path := "/"
	var fs syscall.Statfs_t
	err := syscall.Statfs(path, &fs)
	if err != nil {
		fmt.Println("Error:", err)
	}
	total := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := total - free

	totalConverted := prettyByteSize(total)
	usedConverted := prettyByteSize(used)

	return fmt.Sprintf("%s / %s", usedConverted, totalConverted)
}
