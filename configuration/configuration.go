package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/namsral/flag"
	"gopkg.in/yaml.v2"
)

type TbotConfig struct {
	ApiKey   string `json:"ApiKey" yaml:"apiKey"`
	TargetID int    `json:"TargetID" yaml:"targetID"`
	Active   bool   `json:"Active" yaml:"active"`
	Callback string `json:"Callback" yaml:"callback"`
}

type EmailConfig struct {
	Host     string `json:"Host" yaml:"host"`
	Username string `json:"Username" yaml:"username"`
	Password string `json:"Password" yaml:"password"`
	Port     int    `json:"Port" yaml:"port"`
	Active   bool   `json:"Active" yaml:"active"`
	Callback string `json:"Callback" yaml:"callback"`
}

type Transports struct {
	Tbot  TbotConfig  `json:"tbot" yaml:"tbot"`
	Email EmailConfig `json:"email" yaml:"email"`
}

type Config struct {
	Transports Transports `json:"transports" yaml:"transports"`
}

func Load(configFile string) (config *Config, err error) {
	config = NewConfig()
	switch filepath.Ext(configFile) {
	case ".json":
		if err = LoadJsonConfig(&configFile, config); err != nil {
			return
		}
	case ".yaml":
		if err = LoadYamlConfig(&configFile, config); err != nil {
			return
		}
	case ".env":
		if err = LoadEnvConfig(&configFile, config); err != nil {
			return
		}
	default:
		return nil, fmt.Errorf("invalid format of configuration file")
	}

	if len(os.Args) > 1 && filepath.Ext(os.Args[0]) != ".test" { // тут не нашел лучшего решения исключить тесты из вызова. Буду благодарен, если подскажете варианты
		if err = LoadFromArgs(config); err != nil {
			return
		}
	}

	return
}

func LoadJsonConfig(configFile *string, config *Config) error {
	contents, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Error read config file: %s\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal(contents, config)
	if err != nil {
		return fmt.Errorf("invalid json: %s\n", err)
	}

	return nil
}

func LoadYamlConfig(configFile *string, config *Config) error {
	contents, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Printf("Error read config file: %s\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(contents, config)
	if err != nil {
		return fmt.Errorf("invalid yaml: %s\n", err)
	}

	return nil
}

func LoadEnvConfig(configFile *string, config *Config) error {
	if err := godotenv.Load(*configFile); err != nil {
		return fmt.Errorf("произошла ошибка парсинга файла окружения")
	}

	return nil
}

func LoadFromArgs(config *Config) error {
	flag.Parse()

	return nil
}

func NewConfig() *Config {
	return &Config{}
}
