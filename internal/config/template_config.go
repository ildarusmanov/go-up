package config

import (
	"fmt"
	"os"

	"text/template"

	"gopkg.in/yaml.v3"
)

const (
	DefaultTemplatesConfigPath = "~/.go-up/templates-config.yml"
	TemplatesConfigPathEnv     = "GOUP_TEMPLATES_CONFIG_PATH"
)

type TemplatesConfig struct {
	Templates []*TemplatesConfigItem `yaml:"templates"`
}

type TemplatesConfigItem struct {
	Name             string `yaml:"name"`
	TemplateFilePath string `yaml:"source"`
	ConfigFilePath   string `yaml:"config"`
}

func (cfg *TemplatesConfigItem) GetTemplate(wdir string) (*template.Template, error) {
	path, err := cfg.GetTemplateFilePath(wdir)

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(path)
}

func (cfg *TemplatesConfigItem) GetConfig(wdir string) (*TemplateConfig, error) {
	path, err := cfg.GetTemplateFilePath(wdir)

	if err != nil {
		return nil, err
	}

	return LoadTemplateConfigFile(path)
}

func (cfg *TemplatesConfigItem) GetConfigFilePath(wdir string) (string, error) {
	return getFilePath(wdir, cfg.ConfigFilePath)
}

func (cfg *TemplatesConfigItem) GetTemplateFilePath(wdir string) (string, error) {
	return getFilePath(wdir, cfg.ConfigFilePath)
}

type TemplateConfig struct {
	FileNameTemplate string                 `yaml:"file_name_template"`
	DefaultDirectory string                 `yaml:"default_directory"`
	Vars             map[string]interface{} `yaml:"vars"`
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

func getFilePath(wdir, relativePath string) (string, error) {
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
