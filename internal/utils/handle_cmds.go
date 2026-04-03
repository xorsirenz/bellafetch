package utils

import (
	"flag"
	"fmt"
	"os"
)

func HandleCmd(version string) {
	hFlag := flag.Bool("help", false, "Show help")
	vFlag := flag.Bool("version", false, "Print version")
	flag.BoolVar(hFlag, "h", *hFlag, "alias for -help")
	flag.BoolVar(vFlag, "v", *vFlag, "alias for -version")
	flag.Parse()

	if *vFlag {
		fmt.Println("bellafetch", version)
		os.Exit(0)
	}
	if *hFlag {
		fmt.Println(`Usage:
	bellafetch

Operations:
	bellafetch {-h -help}
	bellafetch {-h -version}`)
		os.Exit(0)
	}
}
