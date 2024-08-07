package cmd

import (
	flag "github.com/spf13/pflag"
)

var configPath string
var projectPath string

func defaultVar(inCmd *flag.FlagSet) {
	inCmd.StringVar(&configPath, "f", "", "config file path")
	inCmd.StringVar(&projectPath, "p", "", "project path")
}

var serviceName string

func selectApiDefaultVar(inCmd *flag.FlagSet) {
	inCmd.StringVar(&serviceName, "s", "", "select api")
}
