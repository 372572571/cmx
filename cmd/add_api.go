package cmd

import (
	config "cmx/v1/logic/aggregate/build_config"
	api_model "cmx/v1/logic/model/api"
	message_model "cmx/v1/logic/model/message"
	"cmx/v1/logic/util"
	"cmx/v1/pkg/logger"
	"context"

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
	Api := api_model.ApiDefinition{
		Definition: make(map[string][]api_model.Api),
	}
	cfg := config.GetDefaultConfig()
	Api.Definition[group] = []api_model.Api{
		{
			Name: funName,
			Http: struct {
				IsPublic  bool   "json:\"is_public\" yaml:\"is_public\""
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
	msg := message_model.Message{
		Definition: make(map[string][]message_model.MessageField),
	}
	msg.Definition[fmt.Sprintf("%s_request", funName)] = []message_model.MessageField{
		{},
	}
	msg.Definition[fmt.Sprintf("%s_response", funName)] = []message_model.MessageField{
		{},
	}
	sb := strings.Builder{}
	apiBuf := util.MustSuccess(yaml.Marshal(&Api))
	sb.Write(apiBuf)
	sb.WriteString("\n")
	msgBuf := util.MustSuccess(yaml.Marshal(&msg))
	sb.Write(msgBuf)
	sb.WriteString("\n")
	// 组是否存在
	apiYaml := cfg.Apis[cfg.SelectApi].ApiYamlPath
	dirPath := path.Join(apiYaml, group)
	if !util.IsHaveDir(dirPath) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			logger.Fatalf(context.Background(),
				"mkdir all error %s", err.Error())
		}
	}
	apiFile := path.Join(dirPath, funName+".yaml")
	if util.IsHaveFile(apiFile) {
		logger.Warnf(context.Background(),
			"file %s is exist", apiFile)
		return
	}
	err := os.WriteFile(
		path.Join(dirPath, funName+".yaml"),
		[]byte(sb.String()), os.ModePerm,
	)
	if err != nil {
		logger.Fatalf(context.Background(),
			"write file error %s", err.Error())
	}
}
