package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.AddConfigPath("./config")
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	err = config.ReadInConfig()
	if err != nil {
		log.Panicln("fatal error config file: %w", err)
	}
	log.Printf("Using config: %s\n", config.ConfigFileUsed())
}

func RelativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
