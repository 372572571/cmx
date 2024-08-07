package config

type TypeConfig struct {
	OutputPath string      `json:"output_path" yaml:"output_path"`
	GoPkgName  string      `json:"go_pkg_name" yaml:"go_pkg_name"`
	ImportPkgs []ImportPkg `json:"import_pkgs" yaml:"import_pkgs"`
}

func (tc TypeConfig) GetImportPaths() []string {
	return getImportPaths(tc.ImportPkgs)
}

func (tc TypeConfig) GetDefaultRefs() []string {
	return getDefaultRefs(tc.ImportPkgs)
}

func (tc TypeConfig) GetOutputPath() string {
	return tc.OutputPath
}

func (tc TypeConfig) GetGoPkgName() string {
	return tc.GoPkgName
}
