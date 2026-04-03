package main

import (
	"github.com/xorsirenz/bellafetch/internal/utils"
	"github.com/xorsirenz/bellafetch/pkg"
)

var version string

func main() {
	config := utils.LoadConfig()
	data := pkg.CheckOS()

	utils.ClearScreen()
	utils.Banner()
	utils.PrintSelectedModules(data, config)
}
