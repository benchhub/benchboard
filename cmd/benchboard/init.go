package main

import (
	"os"

	"github.com/dyweb/gommon/util/fsutil"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create empty .benchboard directory",
	Long:  "Create empty .benchboard directory",
	Run: func(cmd *cobra.Command, args []string) {
		dir := bbProject
		if fsutil.DirExists(dir) {
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
