package config


type ModelConfig struct {
	// model output path
	OutputPath string `json:"output_path" yaml:"output_path"`
	// model package name
	PkgName string `json:"pkg_name" yaml:"pkg_name"`
	// model import packages
	ImportPkgs []ImportPkg `json:"import_pkgs" yaml:"import_pkgs"`
}


func (tc ModelConfig) GetImportPaths() []string {
	return getImportPaths(tc.ImportPkgs)
}

func (tc ModelConfig) GetDefaultRefs() []string {
	return getDefaultRefs(tc.ImportPkgs)
}

func (tc ModelConfig) GetOutputPath() string {
	return tc.OutputPath
}

func (tc ModelConfig) GetPkgName() string {
	return tc.PkgName
}