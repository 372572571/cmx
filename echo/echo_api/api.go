package echo_api

import (
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/util"
	"fmt"
	"strings"
)

type IDefinition interface {
}

type Api struct {
	Group      string
	Definition config.Definition // definition config
	ApiMessage ApiMessage
}

func NewApi(group string, definition config.Definition) *Api {
	return &Api{
		Group:      group,
		Definition: definition,
		ApiMessage: ApiMessage{},
	}
}

func (api *Api) WriteApiGroup(apm []ApiMessage) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("service %s {", util.ToCamelCasing(api.Group)))
	sb.WriteString("\n")
	for _, i := range apm {
		sb.WriteString(i.ApiContent)
		sb.WriteString("\n")
	}
	sb.WriteString("}")
	sb.WriteString("\n")
	api.ApiMessage.IsUnmixed = true
	api.ApiMessage.ApiContent = sb.String()
	request := strings.Builder{}
	response := strings.Builder{}
	for _, i := range apm {
		request.WriteString(i.ApiContentRequest)
		response.WriteString(i.ApiContentResponse)
		if !i.IsUnmixed {
			api.ApiMessage.IsUnmixed = false
		}
	}
	api.ApiMessage.ApiContentRequest = request.String()
	api.ApiMessage.ApiContentResponse = response.String()
	// fmt.Println(api)
}

func (api Api) GetApiParameters() []definition.Api {
	if gourp, found := api.Definition.GetGroup(api.Group); found {
		return gourp
	} else {
		panic(fmt.Sprintf("group %s not found", api.Group))
	}
}

func (api Api) WriteApi(ax definition.Api) string {
	sb := strings.Builder{}
	sb.WriteString(" rpc ")
	sb.WriteString(ax.Name)
	// sb.WriteString(fmt.Sprintf("(%s)", util.ToCamelCasing(ax.Request)))
	sb.WriteString(fmt.Sprintf("(%s)", fmt.Sprintf("%s%s", util.ToCamelCasing(ax.Name), "Request")))
	sb.WriteString("returns")
	// sb.WriteString(fmt.Sprintf("(%s)", util.ToCamelCasing(ax.Response)))
	sb.WriteString(fmt.Sprintf("(%s)", fmt.Sprintf("%s%s", util.ToCamelCasing(ax.Name), "Response")))
	sb.WriteString("{")
	sb.WriteString("option (google.api.http) = {")
	sb.WriteString(fmt.Sprintf("%s: \"/%s\"", ax.Http.Method, ax.Http.Path))
	if ax.Http.Body != "" {
		sb.WriteString(fmt.Sprintf(",body: \"%s\"", ax.Http.Body))
	}
	sb.WriteString("};")
	sb.WriteString("option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) =")
	sb.WriteString("{")
	sb.WriteString(fmt.Sprintf("summary: \"%s\"", ax.Http.Summary))
	description := ax.Description

	if description == "" {
		description = ax.Http.Summary
	}
	if ax.SignType != "" {
		switch ax.SignType {
		case definition.SignTypeHmacSha256Secret:
			description = fmt.Sprintf("%s \\n %s", description,
				"```parsex auth(identity); head(app_id,sign,unix_milli);```")

		}
	}
	description = fmt.Sprintf(",description: \"%s\"", description)
	// description 转义换行符
	sb.WriteString(util.Escape(description))
	sb.WriteString("};")
	sb.WriteString("}")
	return sb.String()
}
