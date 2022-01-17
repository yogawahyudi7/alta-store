package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	Port     int `yaml:"port"`
	Database struct {
		Driver   string `yaml:"driver"`
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		DB_Port  int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	// var defaultConfig AppConfig
	// defaultConfig.Port = 8000
	// defaultConfig.Database.Driver = "mysql"
	// defaultConfig.Database.Address = "localhost"
	// defaultConfig.Database.Port = 3306
	// defaultConfig.Database.Username = "root"
	// defaultConfig.Database.Password = "11223344"
	// defaultConfig.Database.Name = "e_commerces_test"

	// viper.SetConfigType("yaml")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./configs/")

	// if err := viper.ReadInConfig(); err != nil {
	// 	return &defaultConfig
	// }

	// var finalConfig AppConfig
	// err := viper.Unmarshal(&finalConfig)
	// if err != nil {
	// 	log.Info("failed to extract config, will use default value")
	// 	return &defaultConfig
	// }

	// return &finalConfig

	var defaultConfig AppConfig
	defaultConfig.Port = 8080
	defaultConfig.Database.Driver = getEnv("DRIVER", "mysql")
	defaultConfig.Database.Name = getEnv("NAME", "e_commerces")
	defaultConfig.Database.Address = getEnv("ADDRESS", "localhost")
	defaultConfig.Database.DB_Port = 3306
	defaultConfig.Database.Username = getEnv("USERNAME", "root")
	defaultConfig.Database.Password = getEnv("PASSWORD", "11223344")

	return &defaultConfig

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println("=> VALUE", value)
		return value
	}
	return fallback
}
