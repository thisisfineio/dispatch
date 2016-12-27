package main

import (
	"github.com/thisisfineio/gox/goxlib"
	"fmt"
	"os"
	"flag"
	"github.com/thisisfineio/dispatch/dispatchlib"
	"github.com/thisisfineio/variant"
)

func main (){
	paths, err := goxlib.CrossCompile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	flag.Parse()

	// version string overrides version file
	if dispatchlib.VersionString != "" {
		fmt.Println(paths)
	} else {
		// if there's no version file we're not deploying
		if dispatchlib.VersionFile == "" {
			os.Exit(0)
		}

		versions, err := variant.Load(dispatchlib.VersionFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(versions)
	}


}
