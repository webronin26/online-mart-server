package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var c config

type config struct {
	System   SystemConfig
	Database DatabaseConfig
	Facebook FacebookConfig
	Google   GoogleConfig
	Register RegisterConfig
	Socket   SocketConfig
}

// 初始化資料庫
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.Unmarshal(&c)
}

func GetSystemConfig() SystemConfig {
	return c.System
}

func GetDatabaseConfig() DatabaseConfig {
	return c.Database
}

func GetFacebookConfig() FacebookConfig {
	return c.Facebook
}

func GetGoogleConfig() GoogleConfig {
	return c.Google
}

func GetRegisterConfig() RegisterConfig {
	return c.Register
}

func GetSocketConfig() SocketConfig {
	return c.Socket
}
