{{- $model:=. -}} 
{{- $fields:=$model.Fields -}}
{{- $upTableName:=$model.UpTableName}}
{{- $tableName:=$model.TableName}}
package {{getModelPkgName}}

import(
    {{- range $k,$v:=getImportPaths}}
    "{{$v}}"
    {{- end}}
)

{{- range $k,$v:=getDefaultRefs}}
{{- if $v}}
{{$v}}
{{- end}}
{{- end}}

// table comment {{$model.Comment}}
type {{$model.UpTableName}} struct {
    {{- range $k,$v:=sortColumn $model}}
    {{$v.UpFieldName}} {{$v.GoSchema.Type}} `{{$v.GoSchema.Tag}}`
    {{- end}}
}

func (m *{{$upTableName}}) TableName() string {
	return "{{$tableName}}"
}