syntax = "proto3";
{{$api:= .ApiMessage}}
{{printf `package %s;` (getProtoPkgName .)}}

{{printf `option go_package = "%s";` (protoGoPkgName .)}}

{{- range $k,$import := getImportPaths}}
  {{- if $import}}
{{printf `import "%s";` $import}}
  {{- end}}
{{- end}}
{{if eq $api.IsUnmixed false}}
// lib import
  {{- range $k,$import := getOptionsImportPaths}} 
{{printf `import "%s";` $import}}
  {{- end}}
{{- end}}

{{$api.ApiContent}}

{{$api.ApiContentRequest}}

{{$api.ApiContentResponse}}
