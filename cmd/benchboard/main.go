package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/at15/go.ice/ice"
	"github.com/benchhub/benchboard/pkg/common"
	"github.com/benchhub/benchboard/pkg/util/logutil"
)

const (
	myname = "benchboard" // 你的名字
)

var log = logutil.Registry

var (
	app       *ice.App
	bbHome    string
	bbProject string
)

func main() {
	// TODO: serving static assets, keep track of local experiments etc., might support different result format
	app = ice.New(
		ice.Name(myname),
		ice.Description("BenchBoard keep tracks benchmarks you have run and visualize them"),
		ice.Version(common.Version()),
		ice.LogRegistry(log),
	)
	root := ice.NewCmd(app)
	root.AddCommand(globalInitCmd, initCmd, statusCmd, runCmd)
	// TODO: go.ice need to allow extra prerun since update root.PreRun would replace what is applied in go.ice
	bbHome = join(userHome(), common.Dir)
	bbProject = join(cwd(), common.Dir)
	if err := root.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

// https://stackoverflow.com/questions/7922270/obtain-users-home-directory
func userHome() string {
	if u, err := user.Current(); err != nil {
		log.Warnf("can't get current user %v", err)
		return ""
	} else {
		return u.HomeDir
	}
}

func cwd() string {
	if cur, err := os.Getwd(); err != nil {
		log.Warnf("can't get current directory %v", err)
		return ""
	} else {
		return cur
	}
}

func join(p ...string) string {
	return filepath.Join(p...)
}
