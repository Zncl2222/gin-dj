package cmd

import (
	"fmt"

	"github.com/Zncl2222/gin-dj/gin_dj"
	"github.com/spf13/cobra"
)

var createappCmd = &cobra.Command{
	Use:   "createapp [app-name]",
	Short: "Initialize the folder structure of a Gin web app.",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		if !isGoModule() {
			return fmt.Errorf("Current directory is not a go module !")
		}
		return nil
	},
	Long: `This command initializes the folder structure of a Gin web app.`,
	Run:  gin_dj.CreateApp,
}

func init() {
	rootCmd.AddCommand(createappCmd)
}
