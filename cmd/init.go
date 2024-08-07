package cmd

import (
	"cmx/echo/echo_init"
	"cmx/echo/echo_init/data_source"
	"cmx/pkg/config"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var initSteps = []*survey.Select{
	{
		Message: "Confirm initialization",
		Options: []string{"yes", "no"},
		Default: "no",
		Description: func(value string, index int) string {
			if value == "yes" {
				return "yes"
			}
			return "no"
		},
	},
}
var _ = initSteps
var initCmd = &cobra.Command{
	Use: "init",
	// Short: "init project",
	// Long:  `init project`,
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			init string `survey:"is_init"`
		}{}
		err := survey.AskOne(initSteps[0], &answers.init)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if answers.init == "yes" {
			fmt.Println("init")
			echo_init.Generated(config.GetDefaultConfig(), data_source.SourceTypeMysql)
		} else {
			fmt.Println("no init")
		}
	},
}
