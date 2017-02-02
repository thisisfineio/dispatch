package dispatchlib

import (
	"flag"
	"os"
	"github.com/thisisfineio/variant"
	"github.com/google/go-github/github"
	"io/ioutil"
	"errors"
	"github.com/masterzen/azure-sdk-for-go/core/http"
)

var (
	binDir        string
	VersionFile   string
	BumpMajor     bool
	BumpMinor     bool
	VersionString string
	GithubKey     string
	DeployTypes string
	GithubUser string
	Password string
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
	flag.StringVar(&DeployTypes, "d", Github, "A comma separated list of services to deploy to")
	flag.StringVar(&DeployTypes, "deployTypes", Github, "A comma separated list of services to deploy to")
	flag.StringVar(&GithubUser, "githubUser", "", "The Github username to authenticate with for requests that require authentication")
	GithubKey = os.Getenv("GITHUB_API_KEY")
}

type Deployer interface {
	Deploy() error
}

type Config struct {
	Deployers []Deployer
}

type GithubRelease struct {
	Paths []string
	Version *variant.Version
	Description string
	PreRelease bool
}

func NewGithubRelease(paths []string, version *variant.Version, description string, preRelease bool) *GithubRelease {
	return &GithubRelease{paths, version, description, preRelease}
}

func (g *GithubRelease) Deploy() error {
	if GithubUser == "" {
		return errors.New("dispatchlib: must provide githubUser flag")
	}
	if GithubKey == "" {
		return errors.New("dispatchlib: must set GITHUB_API_KEY")
	}
	transport := github.BasicAuthTransport{
		Username: GithubUser,
		Password: GithubKey,
	}
	client := github.NewClient(transport.Client())
	repoService := client.Repositories
	release := &github.RepositoryRelease{}
	release.TagName = &g.Version.VersionString()
	release.Body = &g.Description
	release.Prerelease = &g.PreRelease

	return nil

}