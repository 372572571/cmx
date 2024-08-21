package cmd

import (
	config "cmx/v1/logic/aggregate/build_config"
	"cmx/v1/logic/inside/echo/echo_init"
	"cmx/v1/logic/inside/echo/echo_init/data_source"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var initLocalSteps = []*survey.Select{
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
var initLocalCmd = &cobra.Command{
	Use:   "init_local",
	Short: "init local sql",
	Long:  `init local sql`,
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			init string `survey:"is_init"`
		}{}
		err := survey.AskOne(initLocalSteps[0], &answers.init)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if answers.init == "yes" {
			fmt.Println("init")
			echo_init.Generated(config.GetDefaultConfig(), data_source.SourceTypeLocal)
		} else {
			fmt.Println("no init")
		}
	},
}
