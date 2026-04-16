package main

import (
	"os"

	"github.com/xorsirenz/bellafetch/internal/utils"
	"github.com/xorsirenz/bellafetch/pkg"
)

var version string

func main() {
	if len(os.Args) > 1 {
		utils.HandleCmd(version)
	}

	config := utils.LoadConfig()
	data := pkg.CheckOS()
	ascii := utils.FetchAscii(data.Id, config)

	utils.PrintData(ascii, data, config)
}
