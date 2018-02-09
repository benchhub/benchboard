package main

import (
	"github.com/benchhub/benchboard/pkg/common"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
)

var globalInitCmd = &cobra.Command{
	Use:   "global-init",
	Short: "create user home .benchboard directory",
	Long:  "Create user home .benchboard directory",
	Run: func(cmd *cobra.Command, args []string) {
		// https://stackoverflow.com/questions/7922270/obtain-users-home-directory
		u, err := user.Current()
		if err != nil {
			log.Fatalf("can't get current user %v", err)
			return
		}
		dir := filepath.Join(u.HomeDir, common.Dir)
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
