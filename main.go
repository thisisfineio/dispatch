package main

import (
	_ "github.com/thisisfineio/dispatch/dispatchlib"
	"github.com/thisisfineio/gox/goxlib"
	"fmt"
	"os"
)

func main (){
	paths, err := goxlib.CrossCompile()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(paths)
}
