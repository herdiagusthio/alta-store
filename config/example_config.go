package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type ExampleConfigIPForwarding struct {
	Enabled bool   `mapstructure:"enabled"`
	IP      string `mapstructure:"ip"`
	Port    string `mapstructure:"port"`
}

//AppConfig Application configuration
type ExampleAppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbAddress      string `mapstructure:"db_address"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
}

var examplelock = &sync.Mutex{}
var exampleappConfig *ExampleAppConfig

//GetConfig Initiatilize config in singleton way
func ExampleGetConfig() *ExampleAppConfig {
	if exampleappConfig != nil {
		return exampleappConfig
	}

	examplelock.Lock()
	defer examplelock.Unlock()

	//re-check after locking
	if exampleappConfig != nil {
		return exampleappConfig
	}

	exampleappConfig = exampleinitConfig()

	return exampleappConfig
}

func exampleinitConfig() *ExampleAppConfig {
	var defaultConfig ExampleAppConfig
	var finalConfig ExampleAppConfig

	defaultConfig.AppPort = 8000
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mysql"
	defaultConfig.DbAddress = "127.0.0.1"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "your DB username"
	defaultConfig.DbPassword = "your DB password"
	defaultConfig.DbName = "your DB name"

	//use this if .env file (dont forget to run "source PATH_TO/.env" example "source config/.env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("altastore")
	viper.BindEnv("app_port")
	viper.BindEnv("app_environment")
	viper.BindEnv("db_driver")
	viper.BindEnv("db_address")
	viper.BindEnv("db_port")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
	viper.BindEnv("db_name")
	// viper.ReadInConfig()

	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}
