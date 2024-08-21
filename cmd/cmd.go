package cmd

import (
	"cmx/pkg/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmx-cli",
	Short: "short cmx-cli",
	Long:  `long cmx-cli`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(func() {
		config.InitDefaultConfigYaml(configPath, projectPath)
		config.SetDefaultConfigSelectApi(serviceName)
	})

	defaultVar(rootCmd.PersistentFlags())

	addApiDefaultVar(addApiCmd.PersistentFlags())

	selectApiDefaultVar(apiCmd.PersistentFlags())
	selectApiDefaultVar(initCmd.PersistentFlags())
	selectApiDefaultVar(addApiCmd.PersistentFlags())

	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(repoCmd)
	rootCmd.AddCommand(typesCmd)
	rootCmd.AddCommand(addApiCmd)
	rootCmd.AddCommand(createMsgCmd)
	rootCmd.AddCommand(initLocalCmd)

	rootCmd.AddCommand(versionCmd)
}
