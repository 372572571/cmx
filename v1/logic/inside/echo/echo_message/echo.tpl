syntax = "proto3";
{{- $message:=. }} 

{{printf `package %s;` getProtoPkgName}}

{{printf `option go_package = "%s";` protoGoPkgName}}

{{- range $k,$import := getImportPaths}}
  {{- if $import}}
{{printf `import "%s";` $import}}
  {{- end}}
{{- end}}


{{subMessage $message}}