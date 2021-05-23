package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/secondarykey/yuru"
)

func main() {

	flag.Parse()
	args := flag.Args()

	name := filepath.Join(getHome(), ".yuru.xml")
	if len(args) >= 1 {
		name = args[0]
	}

	err := yuru.Show(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "yuru.Show() error:\n%+v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "bye!")
}

func getHome() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	}
	return os.Getenv(env)
}
