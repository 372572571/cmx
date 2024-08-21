syntax = "proto3";
{{$content:=.Content}}
{{printf `package %s;` (getProtoPkgName)}}

{{printf `option go_package = "%s";` (protoGoPkgName)}}

{{- range $k,$import := getImportPaths}}
  {{- if $import}}
{{printf `import "%s";` $import}}
  {{- end}}
{{- end}}

{{range $idx,$val:=$content}}
{{printf "%s" $val}}
{{end}}