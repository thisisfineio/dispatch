package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thisisfineio/dispatch/dispatchlib"
	"github.com/thisisfineio/gox/goxlib"
)

func main() {
	flag.Usage = func() {
		goxlib.PrintUsage()
		fmt.Println()
		fmt.Println("dispatch specific flags:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
	paths, err := goxlib.CrossCompile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	config, err := dispatchlib.GetConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(config)

	release := dispatchlib.NewGithubRelease(paths, config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = release.Deploy()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
