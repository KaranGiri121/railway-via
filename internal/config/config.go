package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("NOTE_CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the YAML configuration file")

		flag.Parse()
		configPath = *flags
		if configPath == "" {
			log.Fatal("Config path is not set\n use -config `YAML Config Path`")
		}
	}

	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("\nConfig file done not exist %s \n%s", configPath, err.Error())
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	return &cfg
}
