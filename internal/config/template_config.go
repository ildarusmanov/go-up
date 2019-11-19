package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	DefaultTemplatesConfigPath = "~/.go-up/templates-config.yml"
	TemplatesConfigPathEnv     = "GOUP_TEMPLATES_CONFIG_PATH"
)

type TemplateConfigVarValue struct {
	Str      string    `yaml:"str_val"`
	Int      int64     `yaml:"int_val"`
	Float    float64   `yaml:"float_val"`
	StrArr   []string  `yaml:"str_arr_val"`
	IntArr   []int64   `yaml:"int_arr_val"`
	FloatArr []float64 `yaml:"float_arr_val"`
}

type TemplateConfig struct {
	GoupConfig           *GoupConfig                        `yaml:"_"`
	DestinationDirectory string                             `yaml:"destination_directory"`
	BeforeScript         string                             `yaml:"before_script"`
	AfterScript          string                             `yaml:"after_script"`
	Vars                 map[string]*TemplateConfigVarValue `yaml:"vars"`
}

type TemplatesConfig struct {
	Templates []*TemplatesConfigItem `yaml:"templates"`
}

type TemplatesConfigItem struct {
	TemplateConfig *TemplateConfig `yaml:"_"`
	GoupConfig     *GoupConfig     `yaml:"_"`
	Name           string          `yaml:"name"`
	Directory      string          `yaml:"directory"`
}

func (cfgi *TemplatesConfigItem) GetTemplatesDirectory(wdir string) (string, error) {
	return getDirPath(wdir, cfgi.Directory)
}

func (cfgi *TemplatesConfigItem) GetConfigFilePath(wdir string) (string, error) {
	return getFilePath(wdir, cfgi.Directory+"/.template.goup.yml")
}

func (cfgi *TemplatesConfigItem) LoadConfig(wdir string, gcfg *GoupConfig) error {
	path, err := cfgi.GetConfigFilePath(wdir)

	if err != nil {
		return err
	}

	cfg, err := LoadTemplateConfigFile(path)

	if err != nil {
		return err
	}

	cfgi.GoupConfig = gcfg
	cfgi.TemplateConfig = cfg

	return nil
}

func LoadTemplatesConfig() (*TemplatesConfig, error) {
	templatesConfigPath, ok := os.LookupEnv(TemplatesConfigPathEnv)

	if !ok {
		templatesConfigPath = DefaultTemplatesConfigPath
	}

	return loadTemplatesConfigFile(templatesConfigPath)
}

func LoadTemplateConfigFile(configPath string) (*TemplateConfig, error) {
	f, err := os.Open(configPath)

	if err != nil {
		return nil, fmt.Errorf("can not open template config %s: %w\n", configPath, err)
	}

	d := yaml.NewDecoder(f)

	cfg := &TemplateConfig{}

	if err := d.Decode(cfg); err != nil {
		return nil, fmt.Errorf("can not parse template config %s: %w\n", configPath, err)
	}

	return cfg, nil
}

func isFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}

func isDirExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !os.IsNotExist(err)
}

func loadTemplatesConfigFile(templatesConfigPath string) (*TemplatesConfig, error) {
	f, err := os.Open(templatesConfigPath)

	if err != nil {
		return nil, fmt.Errorf("can not open templates config %s: %w\n", templatesConfigPath, err)
	}

	d := yaml.NewDecoder(f)

	cfg := &TemplatesConfig{}

	if err := d.Decode(cfg); err != nil {
		return nil, fmt.Errorf("can not parse templates config %s: %w\n", templatesConfigPath, err)
	}

	return cfg, nil
}

func getDirPath(wdir, relativePath string) (string, error) {
	if isFileExists(wdir + "/" + relativePath) {
		return wdir + "/" + relativePath, nil
	}

	envd, ok := os.LookupEnv(TemplatesConfigPathEnv)

	if ok && isFileExists(envd+"/"+relativePath) {
		return envd + "/" + relativePath, nil
	}

	if isFileExists(DefaultTemplatesConfigPath + "/" + relativePath) {
		return DefaultTemplatesConfigPath + "/" + relativePath, nil
	}

	return "", fmt.Errorf("can not find the file: %s", relativePath)
}

func getFilePath(wdir, relativePath string) (string, error) {
	if isDirExists(wdir + "/" + relativePath) {
		return wdir + "/" + relativePath, nil
	}

	envd, ok := os.LookupEnv(TemplatesConfigPathEnv)

	if ok && isDirExists(envd+"/"+relativePath) {
		return envd + "/" + relativePath, nil
	}

	if isDirExists(DefaultTemplatesConfigPath + "/" + relativePath) {
		return DefaultTemplatesConfigPath + "/" + relativePath, nil
	}

	return "", fmt.Errorf("can not find the file: %s", relativePath)
}
