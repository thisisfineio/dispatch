package dispatchlib

import (
	"flag"
	"os"
)

var (
	binDir        string
	VersionFile   string
	BumpMajor     bool
	BumpMinor     bool
	VersionString string
	GithubKey     string
	DeployType string
)

// supported deploy types
const (
	Github = "github"
)

func init(){
	flag.StringVar(&VersionFile, "versionFile", "", "Specifies the version file to use if uploading to a github release")
	flag.BoolVar(&BumpMajor, "major", false, "Bumps the major version of this release")
	flag.BoolVar(&BumpMinor, "minor", false, "Bumps the minor version of this release")
	flag.StringVar(&VersionString, "version", "", "Specifies an entire version string to use instead of a config file. By default will save the version config as variant.json in the working directory")
	flag.StringVar(&DeployType, "d", Github, "Specifies what services to deploy to. Default is github.")

	GithubKey = os.Getenv("GITHUB_API_KEY")
}

type Deployer interface {
	Deploy() error
}