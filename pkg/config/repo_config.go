package config

type RepoConfig struct {
	// printf "model.%s" .UpCamelCasingTableName }}
	// ModelNameTpl = model
	ModelNameTpl string `json:"model_name_tpl" yaml:"model_name_tpl"`
	ProtoNameTpl string `json:"proto_name_tpl" yaml:"proto_name_tpl"`
	// model output path
	OutputPath string `json:"output_path" yaml:"output_path"`
	// model package name
	PkgName string `json:"pkg_name" yaml:"pkg_name"`
	// model import packages
	ImportPkgs []ImportPkg `json:"import_pkgs" yaml:"import_pkgs"`
	// enable model to proto
	EnableModelToProto bool `json:"enable_model_to_proto" yaml:"enable_model_to_proto"`
}


func (tc RepoConfig) GetImportPaths() []string {
	return getImportPaths(tc.ImportPkgs)
}

func (tc RepoConfig) GetDefaultRefs() []string {
	return getDefaultRefs(tc.ImportPkgs)
}

func (tc RepoConfig) GetOutputPath() string {
	return tc.OutputPath
}

func (tc RepoConfig) GetPkgName() string {
	return tc.PkgName
}
