package dispatchlib

import (
	"os"
	"fmt"
)

var (
	binDir string
)

func init(){

	binDir, _ = os.Getwd()
	fmt.Println(binDir)
}