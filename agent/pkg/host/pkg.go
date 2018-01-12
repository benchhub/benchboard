// Package host collects hardware info and metrics of the host machine
package host

import (
	"github.com/benchhub/benchboard/agent/pkg/util/logutil"
	"fmt"
)

var log = logutil.NewLogger()

type Procfile interface {
	Path() string
}

func init() {
	fmt.Println(log.Level())
	fmt.Println(log.Identity().Package)
	fmt.Println(log.IsInfoEnabled())
	log.Error("init host")
}
