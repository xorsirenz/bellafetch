package utils

import (
	"flag"
	"fmt"
	"os"
)

func HandleCmd(version string) {
	vFlag := flag.Bool("version", false, "Print version")
	hFlag := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *vFlag {
		fmt.Println("bellafetch", version)
		os.Exit(0)
	}
	if *hFlag {
		fmt.Println(`Usage:
	bellafetch

Operations:
	bellafetch {-help}
	bellafetch {-version}`)
		//flag.PrintDefaults()
		os.Exit(0)
	}
}
