package echo_stores

import (
	"cmx/echo/echo_api"
	"cmx/echo/echo_message"
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/util"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/samber/lo"
	"gopkg.in/yaml.v2"
)

type echo struct {
	Content []string
}

//go:embed echo.tpl
var echo_tpl string
var ECHO_TPL *template.Template
var DefaultImportPkg = []config.ImportPkg{
	config.ImportPkg{
		Path: "google/api/annotations.proto",
	},
	config.ImportPkg{
		Path: "google/api/field_behavior.proto",
	},
	config.ImportPkg{
		Path: "protoc-gen-openapiv2/options/annotations.proto",
	},
}

// proto template func
var templateFunc = template.FuncMap{
	"getProtoPkgName": getProtoPkgName,
	"protoGoPkgName":  protoGoPkgName,
	"getImportPaths":  getImportPaths,
}

// output proto pkg name
func getProtoPkgName() string {
	return config.GetDefaultConfig().StoresConfig.ProtoConfig.PkgName
}

// output proto go pkg name
func protoGoPkgName() string {
	return config.GetDefaultConfig().
		StoresConfig.ProtoConfig.GoPkgName
}

// output reference import path
func getImportPaths() []string {
	result := []string{}
	mip := config.GetDefaultConfig().StoresConfig.ProtoConfig.ImportPkgs
	if len(mip) == 0 {
		mip = DefaultImportPkg
	}
	for _, v := range mip {
		result = append(result, v.Path)
	}
	return result
}

func Generated(cfg config.Config) (str string, err error) {
	ECHO_TPL = template.Must(template.New("echo_api").
		Funcs(templateFunc).
		Parse(echo_tpl))
	echo := echo{Content: []string{}}
	echo.Content = append(echo.Content, ApiIncorporateReferences(cfg))
	sb := strings.Builder{}
	err = ECHO_TPL.Execute(&sb, echo)
	if err != nil {
		return
	}
	return sb.String(), nil
}

// incorporate references
// 1. 获取所有api依赖的引用仅限MessageReferenceType=message
func ApiIncorporateReferences(cfg config.Config) string {
	parameters := []string{}
	for k, _ := range cfg.Apis {
		apis := GetSelectApiConf(k, cfg)
		for _, v := range apis {
			api := echo_api.NewApi(v, *config.GetDefaultConfig().GetDefinition())
			apiParameters := api.GetApiParameters()
			for _, parameter := range apiParameters {
				parameters = append(parameters, parameter.Request, parameter.Response)
			}
		}
	}
	var references = []string{}
	// fmt.Printf("cfg.StoresConfig.ForceReferenceFile : %s \n", cfg.StoresConfig.ForceReferenceFile)
	if cfg.StoresConfig.ForceReferenceFile != "" &&
		util.IsHaveFile(cfg.StoresConfig.ForceReferenceFile) {
		buf, err := os.ReadFile(cfg.StoresConfig.ForceReferenceFile)
		util.NoError(err)
		err = yaml.Unmarshal(buf, &references)
		util.NoError(err)
		// fmt.Printf("force reference file: %s \n", cfg.StoresConfig.ForceReferenceFile)
		// fmt.Printf("force reference file: %v \n", references)
	}
	references = lo.Uniq(append(cfg.StoresConfig.ForceReference, references...))
	for _, v := range parameters {
		references = append(references, getMessageReferences(v, cfg)...)
	}
	references = lo.Uniq(references)
	sb := strings.Builder{}
	for _, v := range references {
		message := echo_message.NewMessage(*cfg.GetDefinition(), echo_message.SetDisableIncorporate(true))
		message.InitNode(v)
		message.WriterMessage(&sb)
	}
	return sb.String()
}

func getMessageReferences(parameter string, cfg config.Config) (result []string) {
	route := config.NewReferenceInformation(parameter)
	r := fmt.Sprintf("/%s", route.Route)
	values, found := cfg.GetDefinition().GetMessages(r)
	if !found {
		panic(fmt.Sprintf("not found %s \n", parameter))
	}
	for _, v := range values.Definition[route.Field] {
		if v.Ref.Type != definition.RefTypeMessage {
			continue
		}
		result = append(result, v.Ref.Ref)
		result = append(result, getMessageReferences(v.Ref.Ref, cfg)...)
	}
	return
}

func GetSelectApiConf(selectApi string, cfg config.Config) (apis []string) {
	apiPath := filepath.Join(cfg.ProjectPath,
		selectApi,
		selectApi+".conf")
	apiBuf, err := os.ReadFile(apiPath)
	if err != nil {
		panic(err)
	}
	util.NoError(yaml.Unmarshal(apiBuf, &apis))
	return
}
