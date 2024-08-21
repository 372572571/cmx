package echo_model

import (
	config "cmx/v1/logic/aggregate/build_config"
	"cmx/v1/logic/aggregate/parse"
	_ "embed"
	"go/format"
	"strings"
	"text/template"
)

//go:embed echo_model.tpl
var echo_tpl string
var ECHO_TPL *template.Template

func Generated(m *parse.Model) (b []byte, err error) {

	// model template func
	var templateFunc = template.FuncMap{
		"sortColumn":      sortColumn,
		"getImportPaths":  config.GetDefaultConfig().ModelConfig.GetImportPaths(),
		"getDefaultRefs":  config.GetDefaultConfig().ModelConfig.GetDefaultRefs(),
		"getModelPkgName": config.GetDefaultConfig().ModelConfig.GetPkgName(),
	}

	// 初始化模板
	ECHO_TPL = template.Must(template.New("echo_model").
		Funcs(templateFunc).
		Parse(echo_tpl))
	sb := strings.Builder{}

	err = ECHO_TPL.Execute(&sb, m)
	if err != nil {
		return
	}

	return format.Source([]byte(sb.String()))
}

// sort by idx echo model field
func sortColumn(m *parse.Model) []parse.ModelField {
	return m.SortFields()
}
