package config

import (
	"cmx/pkg/util"
	"encoding/json"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var defaultConfig *Config = &Config{}

func GetDefaultConfig() Config {
	return *defaultConfig
}

func SetDefaultConfigSelectApi(val string) {
	defaultConfig.SelectApi = val
}

type Config struct {
	// project name
	SelectApi string
	// db config
	DBConfig DataBaseConfig `json:"db_config" yaml:"db_config"`
	// targer reference
	TargerReference string `json:"targer_reference" yaml:"targer_reference"`

	// project path
	ProjectPath string `json:"project_path" yaml:"project_path"`

	// if the field supports null enable point
	EnableGoNullPoint bool `json:"enable_go_null_point" yaml:"enable_go_null_point"`

	// if deleted_at field exists, enable gorm soft delete
	EnableGormSoftDelete bool `json:"enable_gorm_soft_delete" yaml:"enable_gorm_soft_delete"`

	// go model enable gorm tag
	EnableGormTag bool `json:"enable_gorm_tag" yaml:"enable_gorm_tag"`
	// proto model enable gorm tag serializer
	EnableGormSerializer bool `json:"enable_gorm_serializer" yaml:"enable_gorm_serializer"`
	// enable bigint to string
	EnableProtoBigIntToString bool `json:"enable_big_int_to_string" yaml:"enable_big_int_to_string"`

	ModelConfig   ModelConfig          `json:"model_config" yaml:"model_config"`
	ProtoConfig   ProtoConfig          `json:"proto_config" yaml:"proto_config"`
	MessageConfig MessageConfig        `json:"message_config" yaml:"message_config"`
	RepoConfig    RepoConfig           `json:"repo_config" yaml:"repo_config"`
	Apis          map[string]ApiConfig `json:"api_config" yaml:"apis"`
	TypeConfig    TypeConfig           `json:"type_config" yaml:"type_config"`
	StoresConfig  StoresConfig         `json:"stores_config" yaml:"stores_config"`

	LinkModle []string `json:"link_modle" yaml:"link_modle"`
	// definition  manager
	definition *Definition `json:"-" yaml:"-"`
}

func (c Config) GetDefinition() *Definition {
	return c.definition
}

type ApiConfig struct {
	IsJoin bool `json:"is_join" yaml:"is_join"`
	// ModelConfig `json:",inline" yaml:",inline"`
	ModelConfig ModelConfig `json:"model_config" yaml:"model_config"`
	ProtoConfig ProtoConfig `json:"proto_config" yaml:"proto_config"`
	ApiYamlPath string      `json:"api_yaml_path" yaml:"api_yaml_path"`
	ApiGoPath   string      `json:"api_go_path" yaml:"api_go_path"`
}

type StoresConfig struct {
	IsEnable           bool        `json:"is_enable" yaml:"is_enable"`
	StoresName         string      `json:"stores_name" yaml:"stores_name"`
	ForceReference     []string    `json:"force_reference" yaml:"force_reference"`
	ForceReferenceFile string      `json:"force_reference_file" yaml:"force_reference_file"`
	ProtoConfig        ProtoConfig `json:"proto_config" yaml:"proto_config"`
}

type DataBaseConfig struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	DbName string `json:"db_name"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
}

type ImportPkg struct {
	Path       string `json:"path" yaml:"path"`
	DefaultRef string `json:"default_ref" yaml:"default_ref"`
}

func InitDefaultConfigYaml(file string, projectDir string) {
	cfg_str, err := os.ReadFile(file)
	if err != nil {
		return
	}
	util.NoError(yaml.Unmarshal([]byte(cfg_str), &defaultConfig))
	if projectDir != "" {
		defaultConfig.ProjectPath = path.Join(projectDir, defaultConfig.ProjectPath)
		defaultConfig.TypeConfig.OutputPath = path.Join(projectDir, defaultConfig.TypeConfig.OutputPath)
		defaultConfig.ModelConfig.OutputPath = path.Join(projectDir, defaultConfig.ModelConfig.OutputPath)
		defaultConfig.StoresConfig.ProtoConfig.OutputPath = path.Join(projectDir, defaultConfig.StoresConfig.ProtoConfig.OutputPath)
		defaultConfig.RepoConfig.OutputPath = path.Join(projectDir, defaultConfig.RepoConfig.OutputPath)
		if defaultConfig.StoresConfig.ForceReferenceFile != "" {
			defaultConfig.StoresConfig.ForceReferenceFile = path.Join(defaultConfig.ProjectPath, defaultConfig.StoresConfig.ForceReferenceFile)
		}
		for k, v := range defaultConfig.Apis {
			_ = v
			api := defaultConfig.Apis[k]
			api.ApiYamlPath = path.Join(projectDir, api.ApiYamlPath)
			api.ProtoConfig.OutputPath = path.Join(projectDir, api.ProtoConfig.OutputPath)
			defaultConfig.Apis[k] = api
		}
	}
	initDefinition(defaultConfig)
}

func InitDefaultConfigJson(file string) {
	cfg_str := util.MustSucc(os.ReadFile(file))
	util.NoError(json.Unmarshal([]byte(cfg_str), &defaultConfig))
	initDefinition(defaultConfig)
}

func InitDefaultConfig(fn func(cfg *Config)) {
	fn(defaultConfig)
	// init definition manager
	initDefinition(defaultConfig)
}

// init definition
func initDefinition(cfg *Config) {
	if cfg.ProjectPath == "" {
		return
	}
	cfg.definition = NewDefinition(cfg.ProjectPath)
	files := util.LoadDirAllFile(cfg.ProjectPath, []string{".yaml"})
	for _, file := range files {
		cfg.definition.AddStatement(file)
		cfg.definition.AddEnums(file)
		cfg.definition.AddTable(file)
		cfg.definition.AddMessage(file)
		cfg.definition.AddApi(file)
	}
}

func getImportPaths(mip []ImportPkg) []string {
	result := []string{}
	for _, v := range mip {
		result = append(result, v.Path)
	}
	return result
}

func getDefaultRefs(mip []ImportPkg) []string {
	result := []string{}
	for _, v := range mip {
		result = append(result, v.DefaultRef)
	}
	return result
}
