package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/thisisfineio/dispatch/dispatchlib"
	"github.com/thisisfineio/gox/goxlib"
	"github.com/thisisfineio/variant"
)

func main (){
	paths, err := goxlib.CrossCompile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	flag.Parse()

	versions, _ := variant.Load(dispatchlib.VersionFile)
	fmt.Println(versions)


	if dispatchlib.GithubKey == "" {
		os.Exit(0)
	}

	// version string overrides version file
	if dispatchlib.VersionString != "" {
		fmt.Println(paths)
	} else {
		// if there's no version file we're not deploying
		if dispatchlib.VersionFile == "" {
			os.Exit(0)
		}


	}

}
