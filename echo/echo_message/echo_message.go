package echo_message

import (
	"cmx/pkg/config"
	_ "embed"
	"strings"
	"text/template"
)

//go:embed echo.tpl
var echo_tpl string
var ECHO_TPL *template.Template

func Generated(referenceString string) (b []byte, err error) {
	var templateFunc = template.FuncMap{
		"getProtoPkgName": config.GetDefaultConfig().MessageConfig.GetPkgName,
		"protoGoPkgName":  config.GetDefaultConfig().MessageConfig.GetGoPkgName,
		"getImportPaths":  config.GetDefaultConfig().MessageConfig.GetImportPaths,
		"subMessage":      subMessage,
	}
	ECHO_TPL = template.Must(template.New("echo_message").
		Funcs(templateFunc).
		Parse(echo_tpl))

	ms := NewMessage(*config.GetDefaultConfig().GetDefinition())
	ms.InitNode(referenceString)

	sb := strings.Builder{}

	err = ECHO_TPL.Execute(&sb, ms)
	if err != nil {
		return
	}
	return []byte(sb.String()), nil
}

// proto template func

func subMessage(ms Message) string {
	rw := strings.Builder{}
	ms.WriterMessage(&rw)
	return rw.String()
}
