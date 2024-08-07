package cmd

import (
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/util"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

var funcName string
var groupName string

func addApiDefaultVar(inCmd *flag.FlagSet) {
	inCmd.StringVar(&funcName, "fn", "", "function name")
	inCmd.StringVar(&groupName, "gn", "", "group name")
}

var addApiCmd = &cobra.Command{
	Use:   "add-api",
	Short: "add api",
	Long:  `add api`,
	Run: func(cmd *cobra.Command, args []string) {
		if funcName == "" || groupName == "" {
			flag.Usage()
			return
		}
		addApi(groupName, funcName)
	},
}

// createmodel -a addapi -gn register -fn enterprise -f build.yml
func addApi(group string, funName string) {
	Api := definition.Apidefinition{
		Definition: make(map[string][]definition.Api),
	}
	cfg := config.GetDefaultConfig()
	Api.Definition[group] = []definition.Api{
		definition.Api{
			Name: funName,
			Http: struct {
				IsOpenApi bool   "json:\"is_open_api\" yaml:\"is_open_api\""
				Method    string "json:\"method\" yaml:\"method\""
				Path      string "json:\"path\" yaml:\"path\""
				Body      string "json:\"body\" yaml:\"body\""
				Summary   string "json:\"summary\" yaml:\"summary\""
			}{
				Method: "post",
				Body:   "*",
				Path: fmt.Sprintf("%s/%s", strings.ReplaceAll(group, "_", "-"),
					strings.ReplaceAll(funName, "_", "-")),
			},
			Request:  fmt.Sprintf("%s.%s.%s.%s_request", cfg.SelectApi, group, funName, funName),
			Response: fmt.Sprintf("%s.%s.%s.%s_response", cfg.SelectApi, group, funName, funName),
		},
	}
	msg := definition.MessageDefinition{
		Definition: make(map[string][]definition.MessageField),
	}
	msg.Definition[fmt.Sprintf("%s_request", funName)] = []definition.MessageField{
		{},
	}
	msg.Definition[fmt.Sprintf("%s_response", funName)] = []definition.MessageField{
		{},
	}
	sb := strings.Builder{}
	apiBuf, err := yaml.Marshal(&Api)
	if err != nil {
		panic(err)
	}
	sb.Write(apiBuf)
	sb.WriteString("\n")
	msgBuf, err := yaml.Marshal(&msg)
	if err != nil {
		panic(err)
	}
	sb.Write(msgBuf)
	sb.WriteString("\n")
	// 组是否存在
	apiYaml := cfg.Apis[cfg.SelectApi].ApiYamlPath
	dirPath := path.Join(apiYaml, group)
	if !util.IsHaveDir(dirPath) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	apiFile := path.Join(dirPath, funName+".yaml")
	if util.IsHaveFile(apiFile) {
		fmt.Printf("file is exists %s\n", apiFile)
		return
	}
	err = os.WriteFile(
		path.Join(dirPath, funName+".yaml"),
		[]byte(sb.String()), os.ModePerm,
	)
	if err != nil {
		panic(err)
	}
}
