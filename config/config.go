package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	Mysql MysqlConfig
	AWS   AWSConfig
	Redis RedisConfig
}

type AppConfig struct {
	Version      string
	Port         string
	Mode         string
	Secret       string
	MigrationURL string
}
type MysqlConfig struct {
	Host          string
	ContainerName string
	Port          string
	User          string
	Password      string
	DBName        string
}
type AWSConfig struct {
	Region    string
	APIKey    string
	SecretKey string
	S3Bucket  string
	S3Domain  string
}
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	v.AddConfigPath("config")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		// check is not found file config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	var c Config // Unmarshal data config have get in file config then get into c
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
