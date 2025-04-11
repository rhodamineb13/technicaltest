package config

import "github.com/spf13/viper"

type EnvConfig struct {
	DatabaseName       string `mapstructure:"DB_NAME"`
	DatabasePort       string `mapstructure:"DB_PORT"`
	DatabaseUser       string `mapstructure:"DB_USER"`
	DatabasePass       string `mapstructure:"DB_PASS"`
	DatabaseHost       string `mapstructure:"DB_HOST"`
	RedisPort          string `mapstructure:"REDIS_PORT"`
	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSAccessKeyID     string `mapstructure:"AWS_AKID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SAK"`
	AWSBucketName      string `mapstructure:"AWS_BUCK"`
}

var Conf *EnvConfig

func NewConfig() {
	var conf *EnvConfig

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	Conf = conf
}
