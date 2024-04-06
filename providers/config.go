package providers

import (
	"flag"
	"os"

	envconfig "github.com/kelseyhightower/envconfig"
	yaml "gopkg.in/yaml.v3"
)

type OpenAiConfig struct {
	SecretKey string `yaml:"secret"`
}
type HttpConfig struct {
	Address string `yaml:"address"`
}
type AppConfig struct {
	OpenAiConfig OpenAiConfig `yaml:"openAi"`
	HttpConfig   HttpConfig   `yaml:"http"`
}

var (
	config     *AppConfig
	configPath string
)

func loadConfig() AppConfig {
	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var cnfg AppConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cnfg)
	if err != nil {
		panic(err)
	}
	err = envconfig.Process("", &cnfg)
	if err != nil {
		panic(err)
	}
	return cnfg
}
func GetConfig(path string) (AppConfig, error) {
	if config == nil {
		flag.StringVar(&configPath, "config", path, "path to config file")
		flag.Parse()
		cnf := loadConfig()
		config = &cnf
	}
	return *config, nil
}
