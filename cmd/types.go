package cmd

import (
	config "cmx/v1/logic/aggregate/build_config"
	"cmx/v1/logic/inside/echo/echo_type"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
	// flag "github.com/spf13/pflag"
)

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "create types",
	Long:  `create types`,
	Run: func(cmd *cobra.Command, args []string) {
		types()
	},
}

// types generate types file
func types() {
	outPath := filepath.Join(config.GetDefaultConfig().TypeConfig.OutputPath)
	os.MkdirAll(outPath, os.ModePerm)
	content, err := echo_type.Generated(config.GetDefaultConfig())
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(
		path.Join(outPath, "types.go"),
		[]byte(content), os.ModePerm,
	)
	if err != nil {
		panic(err)
	}
}
