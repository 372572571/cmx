package config

type ProtoConfig struct {
	// proto output path
	OutputPath string `json:"output_path" yaml:"output_path"`
	// proto package name
	PkgName string `json:"pkg_name" yaml:"pkg_name"`
	// proto to go package name
	GoPkgName string `json:"go_pkg_name" yaml:"go_pkg_name"`
	// proto import packages
	ImportPkgs        []ImportPkg `json:"import_pkgs" yaml:"import_pkgs"`
	// proto options import packages
	OptionsImportPkgs []ImportPkg `json:"options_import_pkgs" yaml:"options_import_pkgs"`
}

func (tc ProtoConfig) GetImportPaths() []string {
	return getImportPaths(tc.ImportPkgs)
}

func (tc ProtoConfig) GetDefaultRefs() []string {
	return getDefaultRefs(tc.ImportPkgs)
}

func (tc ProtoConfig) GetOutputPath() string {
	return tc.OutputPath
}

func (tc ProtoConfig) GetGoPkgName() string {
	return tc.GoPkgName
}

func (tc ProtoConfig) GetPkgName() string {
	return tc.PkgName
}

func (tc ProtoConfig) GetOptionsImportPaths() []string {
	return getImportPaths(tc.OptionsImportPkgs)
}