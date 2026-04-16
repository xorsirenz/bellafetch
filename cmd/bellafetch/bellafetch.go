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

	config, _ := utils.LoadConfig()
	data := pkg.CheckOS()

	utils.ClearScreen()
	utils.Banner()
	selectedModules := utils.BuildSelectedModules(data, config)
	utils.RenderAsciiWithSelectedModules(utils.NoAscii(), selectedModules)
} 
