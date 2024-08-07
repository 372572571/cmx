syntax = "proto3";
{{- $model:=. -}} 
{{- $fields:=$model.Fields -}}
{{- $upTableName:=$model.UpTableName}}
{{- $tableName:=$model.TableName}}

{{printf `package %s;` getProtoPkgName}}

{{printf `option go_package = "%s";` protoGoPkgName}}

{{- range $k,$import := getImportPaths}}
  {{- if $import}}
{{printf `import "%s" ;` $import}}
  {{- end}}
{{- end}}

{{/* compatible model */}}
message {{$upTableName}} {
  option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {
    json_schema: {
      title: "{{$upTableName}}"
      description: "{{$model.Comment}}"}
  };
{{sortColumnByFieldString $model}}
}

message {{$upTableName}}Array {
  repeated {{$upTableName}} list = 1;
}