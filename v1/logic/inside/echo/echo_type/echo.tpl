package {{getPkgName}}

import(
    "strconv"
    {{- range $k,$v:=getImportPaths}}
    "{{$v}}"
    {{- end}}
)

{{- range $k,$v:=getDefaultRefs}}
{{- if $v}}
{{$v}}
{{- end}}
{{- end}}

{{.}}