package cmd

import (
	"github.com/Zncl2222/gin-dj/gin_dj"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Gin-based web application template",
	Args:  cobra.MinimumNArgs(1),
	Run:   gin_dj.ProjectInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
