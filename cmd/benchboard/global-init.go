package main

import (
	"os"

	"github.com/spf13/cobra"
)

var globalInitCmd = &cobra.Command{
	Use:   "global-init",
	Short: "create user home .benchboard directory",
	Long:  "Create user home .benchboard directory",
	Run: func(cmd *cobra.Command, args []string) {
		dir := bbHome
		if dirExists(dir) {
			log.Infof("%s already exists", dir)
			return
		}
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			log.Fatalf("failed to create %s", err)
			return
		}
		// TODO: create files based on template
		log.Infof("global benchboard initialized to %s", dir)
	},
}
