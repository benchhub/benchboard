package main

import (
	"fmt"
	"os"

	"github.com/at15/go.ice/ice"
	"github.com/benchhub/benchboard/pkg/common"
	"github.com/benchhub/benchboard/pkg/util/logutil"
)

const (
	myname = "benchboard" // 你的名字
)

var app *ice.App
var log = logutil.Registry

func main() {
	// TODO: serving static assets, keep track of local experiments etc., might support different result format
	app = ice.New(
		ice.Name(myname),
		ice.Description("BenchBoard keep tracks benchmarks you have run and visualize them"),
		ice.Version(common.Version()),
		ice.LogRegistry(log),
	)
	root := ice.NewCmd(app)
	if err := root.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
