package main

import (
	"github.com/spf13/cobra"

	"github.com/benchhub/benchboard/pkg/common"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run specific benchmark",
	Long:  "Run specific benchmark based on config",
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("run %s", join(bbProject, common.ConfigFile))
		acquireLockOrExit()
		releaseLock()
	},
}
