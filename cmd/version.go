package cmd

import "github.com/spf13/cobra"

var versionCmd = &cobra.Command{
	Use:   "version",
	Long:  `version`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("test-v0.0.3")
	},
}
