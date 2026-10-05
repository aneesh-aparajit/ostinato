package config

import (
	"fmt"
	"strings"

	"github.com/aneesh-aparajit/ostinato/pkg/logger"
	"github.com/spf13/viper"
)

const DefaultPath = "resources/properties.yaml"

type Config struct {
	App    string        `mapstructure:"app"`
	Logger logger.Config `mapstructure:"logger"`
}

func Load(path string) (Config, error) {
	v := viper.New()

	v.SetEnvPrefix("Ostinato")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg Config
	if err := v.UnmarshalExact(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config %s: %w", path, err)
	}
	return cfg, nil
}
