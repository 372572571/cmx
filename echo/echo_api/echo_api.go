package echo_api

import (
	"cmx/echo/echo_message"
	"cmx/pkg/config"
	_ "embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/samber/lo"
)

//go:embed echo.tpl
var echo_tpl string

//go:embed echov2.tpl
var echov2_tpl string

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

func GeneratedGroup(group string) (string, error) {
	sb := strings.Builder{}
	_ = sb
	api := NewApi(group, *config.GetDefaultConfig().GetDefinition())
	standAlone := writeStandAloneApiV2(*api)
	api.WriteApiGroup(lo.MapToSlice(standAlone,
		func(key string, item ApiMessage) ApiMessage {
			return item
		}))
	// fmt.Println(api)
	ECHO_TPL = template.Must(template.New("echo_api").
		Funcs(templateFunc).
		Parse(echo_tpl))
	nsb := strings.Builder{}
	err := ECHO_TPL.Execute(&nsb, api)
	return nsb.String(), err
}

func Generated(group string) (b map[string]string, err error) {
	b = map[string]string{}
	ECHO_TPL = template.Must(template.New("echo_api").
		Funcs(templateFunc).
		Parse(echov2_tpl))

	api := NewApi(group, *config.GetDefaultConfig().GetDefinition())
	standAlone := writeStandAloneApiV2(*api)
	for k, v := range standAlone {
		sb := strings.Builder{}
		api.ApiMessage = v
		err = ECHO_TPL.Execute(&sb, api)
		if err != nil {
			return b, err
		}
		b[k] = fmt.Sprintf("%s\n%s", b[k], sb.String())
	}

	return b, err
}

// proto template func
var templateFunc = template.FuncMap{
	"getProtoPkgName":       getProtoPkgName,
	"protoGoPkgName":        protoGoPkgName,
	"getImportPaths":        getImportPaths,
	"getOptionsImportPaths": getOptionsImportPaths,
}

// output proto pkg name
func getProtoPkgName(api Api) string {
	cfg := config.GetDefaultConfig()
	// cfg.Apis[cfg.SelectApi]
	return strings.ReplaceAll(cfg.Apis[cfg.SelectApi].ProtoConfig.PkgName, "${group}", api.Group)
}

// output proto go pkg name
func protoGoPkgName(api Api) string {
	cfg := config.GetDefaultConfig()
	return strings.ReplaceAll(
		cfg.Apis[cfg.SelectApi].ProtoConfig.GoPkgName, "${group}", api.Group)
}

// output reference import path
func getImportPaths() []string {
	result := []string{}
	cfg := config.GetDefaultConfig()
	mip := cfg.Apis[cfg.SelectApi].ProtoConfig.ImportPkgs
	if len(mip) == 0 {
		mip = DefaultImportPkg
	}
	for _, v := range mip {
		result = append(result, v.Path)
	}
	return result
}

func getOptionsImportPaths() []string {
	cfg := config.GetDefaultConfig()
	return cfg.Apis[cfg.SelectApi].ProtoConfig.GetOptionsImportPaths()
}

func writeStandAloneApiV2(api Api) map[string]ApiMessage {
	group := api.GetApiParameters()
	result := map[string]ApiMessage{}
	cfg := config.GetDefaultConfig()
	for _, v := range group {
		// 防止重名问题(临时解决)
		if !strings.HasPrefix(v.Request, cfg.SelectApi+".") || !strings.HasPrefix(v.Response, cfg.SelectApi+".") {
			continue
		}
		if _, ok := result[v.Name]; ok {
			continue
		}
		itemApiMessage := ApiMessage{}
		apiSb := strings.Builder{}
		definition := config.GetDefaultConfig().GetDefinition()
		// apiSb.WriteString("//")
		apiSb.WriteString(api.WriteApi(v))
		// apiSb.WriteString("*/")
		itemApiMessage.ApiContent = apiSb.String()

		requestSb := strings.Builder{}
		request := echo_message.NewMessage(
			*definition,
			echo_message.SetDisableIncorporate(cfg.StoresConfig.IsEnable),
			echo_message.SetIncorporatePkName(cfg.StoresConfig.StoresName),
		)

		request.InitNode(v.Request)
		request.WriterMessage(&requestSb)
		itemApiMessage.ApiContentRequest = requestSb.String()

		responseSb := strings.Builder{}
		response := echo_message.NewMessage(
			*definition,
			echo_message.SetDisableIncorporate(cfg.StoresConfig.IsEnable),
			echo_message.SetIncorporatePkName(cfg.StoresConfig.StoresName),
		)
		response.InitNode(v.Response)
		response.WriterMessage(&responseSb)

		if cfg.StoresConfig.IsEnable {
			openRequestSb := strings.Builder{}
			openResponseSb := strings.Builder{}
			openRequest := echo_message.NewMessage(
				*definition,
				echo_message.SetDisableIncorporate(!cfg.StoresConfig.IsEnable),
			)
			openResponse := echo_message.NewMessage(
				*definition,
				echo_message.SetDisableIncorporate(!cfg.StoresConfig.IsEnable),
			)
			openRequest.InitNode(v.Request)
			openRequest.WriterMessage(&openRequestSb)
			openResponse.InitNode(v.Response)
			openResponse.WriterMessage(&openResponseSb)
			itemApiMessage.OpenContentRequest = openRequestSb.String()
			itemApiMessage.OpenContentResponse = openResponseSb.String()
		}
		itemApiMessage.IsCommentedOutputOnly = v.Http.IsOpenApi
		itemApiMessage.ApiContentResponse = responseSb.String()
		if response.IsWriterSubMessage() || request.IsWriterSubMessage() {
			itemApiMessage.IsUnmixed = false
		} else {
			itemApiMessage.IsUnmixed = true
		}
		result[v.Name] = itemApiMessage
	}
	return result
}
