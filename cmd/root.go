package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gin-dj [init|createapp]",
	Short: "A command line tool for Gin-based web application templates",
	Long:  `gin-dj is a command line tool that helps create and manage Gin-based web application templates, similar to the Django framework.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to gin-dj!")
		fmt.Println("gin-dj is a command line tool that helps create and manage Gin-based web application templates, similar to the Django framework.")
		fmt.Println("Use `gin-dj init [projectname]` to initialize the project, and use `gin-dj createapp [appname]` to create the app.")
	},
}

func Execute() {
	rootCmd.Execute()
}
