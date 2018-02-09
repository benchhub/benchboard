package main

import (
	"os"
	"path/filepath"

	"github.com/benchhub/benchboard/pkg/common"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create empty .benchboard directory",
	Long:  "Create empty .benchboard directory",
	Run: func(cmd *cobra.Command, args []string) {
		cur, err := os.Getwd()
		if err != nil {
			log.Fatalf("can't get current directory %v", err)
			return
		}
		dir := filepath.Join(cur, common.Dir)
		if dirExists(dir) {
			log.Infof("%s already exists", dir)
			return
		}
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			log.Fatalf("failed to create %s", err)
			return
		}
		// TODO: create file based on templates etc.
		log.Infof("local benchboard initialized to %s", dir)
	},
}

func dirExists(path string) bool {
	if info, err := os.Stat(path); err == nil {
		return info.IsDir()
	}
	return false
}
