package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration holds data necessery for configuring application
type Configuration struct {
	Server Server      `yaml:"server,omitempty"`
	DB     Database    `yaml:"database,omitempty"`
	JWT    JWT         `yaml:"jwt,omitempty"`
	App    Application `yaml:"application,omitempty"`
}

// Server holds data necessery for server configuration
type Server struct {
	Port         string `yaml:"port,omitempty"`
	Debug        bool   `yaml:"debug,omitempty"`
	ReadTimeout  int    `yaml:"read_timeout_seconds,omitempty"`
	WriteTimeout int    `yaml:"write_timeout_seconds,omitempty"`
}

// Database holds data necessery for database configuration
type Database struct {
	DSN        string `yaml:"dsn,omitempty"`
	LogQueries bool   `yaml:"log_queries,omitempty"`
	Timeout    int    `yaml:"timeout_seconds,omitempty"`
}

// JWT holds data necessery for JWT configuration
type JWT struct {
	Secret           string `yaml:"secret,omitempty"`
	Duration         int    `yaml:"duration_minutes,omitempty"`
	RefreshDuration  int    `yaml:"refresh_duration_minutes,omitempty"`
	MaxRefresh       int    `yaml:"max_refresh_minutes,omitempty"`
	SigningAlgorithm string `yaml:"signing_algorithm,omitempty"`
}

// Application holds application configuration details
type Application struct {
	MinPasswordStr int    `yaml:"min_password_strength,omitempty"`
	SwaggerUIPath  string `yaml:"swagger_ui_path,omitempty"`
	ErrorFile      string `yaml:"error_file,omitempty"`
}

// LoadConfig - ...
func LoadConfig(configPaths ...string) (*Configuration, error) {

	cfg := &Configuration{}

	v := viper.New()

	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("arya")

	v.AutomaticEnv()

	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("port", 8080)
	v.SetDefault("signing_algorithm", "HS256")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("Could not unmarshal into `config`")
	}

	// TODO: Use Validation of Struct - (i.e. ozzo-validation)

	return cfg, nil
}
