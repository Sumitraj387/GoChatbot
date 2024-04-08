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
	Address string `yaml:"address" envconfig:"LISTEN_ADDRESS"`
}
type AppConfig struct {
	OpenAiConfig OpenAiConfig `yaml:"openAi"`
	HttpConfig   HttpConfig   `yaml:"http"`
	DbConfig     DbConfig     `yaml:"db"`
}
type DbConfig struct {
	Host        string `yaml:"host" envconfig:"DB_HOST"`
	Port        int    `yaml:"port" envconfig:"DB_PORT"`
	User        string `yaml:"user" envconfig:"DB_USER"`
	Password    string `yaml:"password" envconfig:"DB_PASSWORD"`
	Name        string `yaml:"name" envconfig:"DB_NAME"`
	Timeout     int    `yaml:"timeout" envconfig:"DB_CONNECTION_TIMEOUT"`
	MaxOpenConn int    `yaml:"maxOpenConn" envconfig:"DB_MAX_OPEN_CONNECTIONS"`
	MaxIdleConn int    `yaml:"maxIdleConn" envconfig:"DB_MAX_IDLE_CONNECTIONS"`
	SearchPath  string `yaml:"searchPath" envconfig:"DB_SEARCH_PATH"`
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
