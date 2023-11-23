package config

import (
	"strings"

	"github.com/spf13/viper"
)

var (
	defaultConfig = map[string]interface{}{
		"server.addr":                ":8000",
		"server.mode":                "",
		"server.pagesize":            12,
		"logger.out":                 "stdout",
		"logger.level":               "info",
		"logger.file_format":         "%Y%m%d",
		"logger.file_rotation_count": 3,
		"database.dsn":               "sqlite://poetry.db",
		"database.echo":              false,
	}
)

type Config struct {
	*viper.Viper
}

func Default() *Config {
	conf := &Config{
		Viper: viper.New(),
	}
	configs := []map[string]interface{}{
		defaultConfig,
	}
	for _, config := range configs {
		for k, v := range config {
			conf.SetDefault(k, v)
		}
	}

	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	conf.SetEnvPrefix("poetry")
	conf.BindEnv("server.addr")
	conf.BindEnv("database.dsn")
	return conf
}
