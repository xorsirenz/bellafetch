package main

import (
	"fmt"

	"github.com/xorsirenz/bellafetch/internal/utils"
	"github.com/xorsirenz/bellafetch/pkg"
)

var version string

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println("Error: Cannot load config file.")
	}
	data := pkg.CheckOS()

	utils.ClearScreen()
	utils.Banner()
	utils.PrintSelectedModules(data, config)
}
