package main

import "github.com/spf13/cobra"

// print project name, recent benchmarks, kind of git status
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status of current project",
	Long:  "Print status of current project",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("need to print status")
	},
}
