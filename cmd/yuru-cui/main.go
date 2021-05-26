package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/secondarykey/yuru"
)

func init() {
}

func main() {

	flag.Parse()
	args := flag.Args()

	name := ""
	if len(args) >= 1 {
		name = args[0]
	}

	err := yuru.Print(name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "yuru print error:\n%+v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Success")
}
