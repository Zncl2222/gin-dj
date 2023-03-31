package main

import (
	"fmt"
	"os"

	"github.com/Zncl2222/gin-dj/gin_dj"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-dj",
	Short: "A command line tool for Gin-based web application templates",
	Long:  `gin-dj is a command line tool that helps create and manage Gin-based web application templates, similar to the Django framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gin-dj!")
	},
}

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Gin-based web application template",
	Args:  cobra.MinimumNArgs(1),
	Run:   gin_dj.TemplateInit,
}

var initginCmd = &cobra.Command{
	Use:   "createapp",
	Short: "Initialize the folder structure of a Gin web app.",
	Long:  `This command initializes the folder structure of a Gin web app.`,
	Run:   gin_dj.CreateApp,
}

func main() {
	rootCmd.AddCommand(initCmd, initginCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
