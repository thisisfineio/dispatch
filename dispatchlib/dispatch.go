package dispatchlib

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"reflect"

	"github.com/google/go-github/github"
	"github.com/thisisfineio/variant"
)

var (
	VersionString string
	GithubKey     string
	GithubUser    string
	Repo          string
	Password      string
	Owner         string
	Target        string
	TagName       string
	configPath    string
	Description   string
	PreRelease    bool
	config        *Config
)

// supported deploy types
const (
	Github = "github"
)

func init() {
	flag.StringVar(&GithubUser, "githubUser", "", "The Github username to authenticate with for requests that require authentication")
	flag.StringVar(&Repo, "repo", "", "The repository to create a release in")
	flag.StringVar(&Owner, "owner", "", "The owner of the repository")
	flag.StringVar(&Target, "target", "", "The branch target (default is master) ")
	flag.StringVar(&Password, "password", "", "The password to use for github authentication")
	flag.StringVar(&TagName, "tag", "", "The tag to use for this release")
	flag.StringVar(&Description, "description", "", "The description of this release")
	flag.StringVar(&configPath, "c", "", "The path to a configuration file to use for this release")
	flag.BoolVar(&PreRelease, "pre-release", false, "Whether or not this release is a pre release")
}

func GetConfig() (*Config, error) {
	c := &Config{}
	if configPath != "" {
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, c)
		if err != nil {
			return nil, err
		}
	}
	if GithubUser != "" {
		c.User = GithubUser
	}
	if Repo != "" {
		c.Repo = Repo
	}
	if Owner != "" {
		c.Owner = Owner
	}
	if Target != "" {
		c.Target = Target
	}
	if Password != "" {
		c.Password = Password
	}
	if TagName != "" {
		c.TagName = TagName
	}
	if Description != "" {
		c.Description = Description
	}
	if PreRelease {
		c.PreRelease = true
	}
}

type Deployer interface {
	Deploy() error
}

type Config struct {
	TagName     string
	Description string
	Target      string
	Title       string
	Owner       string
	Repo        string
	User        string
	Password    string
	PreRelease  bool
}

func (c *Config) Validate() error {
	v := reflect.ValueOf(c).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			s := f.Interface().(string)
			if s == "" {
				return fmt.Errorf("You must provide %s", v.Type().Field(i).Name)
			}
		}
	}
}

type GithubRelease struct {
	Paths       []string
	c *Config
}

func Deploy(d Deployer) error {
	return d.Deploy()
}

func NewGithubRelease(paths []string, c *Config) *GithubRelease {
	return &GithubRelease{paths, c}
}

func (g *GithubRelease) Deploy() error {

	transport := github.BasicAuthTransport{
		Username: g.c.User,
		Password: g.c.Password,
	}
	client := github.NewClient(transport.Client())
	repoService := client.Repositories
	release := &github.RepositoryRelease{}
	release.TagName = &g.c.TagName
	release.Body = &g.c.Description
	release.Prerelease = &g.c.PreRelease
	release.TargetCommitish = &g.c.Target
	release.Name = &g.c.Title
	rel, _, err := repoService.CreateRelease(g.c.Owner, g.c.Repo, release)
	if err != nil {
		return err
	}
	for _, p := range g.Paths {
		f, err := os.Open(p)
		defer f.Close()
		if err != nil {
			return err
		}
		asset, _, err := repoService.UploadReleaseAsset(Owner, Repo, *rel.ID, nil, f)
		if err != nil {
			return err
		}
		fmt.Printf("Asset %s uploaded and available at %s", p, *asset.BrowserDownloadURL)
	}
	return nil

}
