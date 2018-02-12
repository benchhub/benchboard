package main

import (
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create empty .benchboard directory",
	Long:  "Create empty .benchboard directory",
	Run: func(cmd *cobra.Command, args []string) {
		dir := bbProject
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
