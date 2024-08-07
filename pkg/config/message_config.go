package config

type MessageConfig struct {
	// proto output path
	OutputPath string `json:"output_path" yaml:"output_path"`
	// proto package name
	PkgName string `json:"pkg_name" yaml:"pkg_name"`
	// proto to go package name
	GoPkgName string `json:"go_pkg_name" yaml:"go_pkg_name"`
	// proto import packages
	ImportPkgs []ImportPkg `json:"import_pkgs" yaml:"import_pkgs"`
}

func (tc MessageConfig) GetImportPaths() []string {
	return getImportPaths(tc.ImportPkgs)
}

func (tc MessageConfig) GetDefaultRefs() []string {
	return getDefaultRefs(tc.ImportPkgs)
}

func (tc MessageConfig) GetOutputPath() string {
	return tc.OutputPath
}

func (tc MessageConfig) GetGoPkgName() string {
	return tc.GoPkgName
}

func (tc MessageConfig) GetPkgName() string {
	return tc.PkgName
}
