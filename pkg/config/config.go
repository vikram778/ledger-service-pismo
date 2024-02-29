package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

// Config struct holds the application config required
type Config struct {
	Postgres PostgresConfig
}

// PostgresConfig config is the connection config required to set up application connection with postgres
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  string
	PgDriver           string
}

// LoadConfig takes in the file path from which config needs to read from
// Param - filename type string is the file path from which the config needs to be read from
// Returns - viper object
// Returns - error if any
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigFile(filename)
	v.SetConfigType("yml")
	v.AddConfigPath(filename)
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig parses config from viper and assign it to config instance
// Param - pointer object to viper instance
// Returns - pointer instance of config object
// Returns - error if any
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

// GetConfig reads the config file path passed and load config from the file
// Param - configPath type string is the config file path
// Returns - pointer instance of config model
// Returns - error if any
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// GetConfigPath returns the config path of the file to be used for application
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "config/config-docker.yml"
	}
	return "config/config-local.yml"
}
