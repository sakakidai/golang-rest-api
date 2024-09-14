package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	AppName           string `mapstructure:"app_name"`
	DebugMode         bool   `mapstructure:"debug_mode"`
	EmailVerification bool   `mapstructure:"email_verification"`
	SMTP              struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"smtp"`
}

var (
	config Config
	once   sync.Once
)

func GetConfig() Config {
	return config
}

func LoadConfig() {
	once.Do(func() {
		fmt.Println("Load config")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")

		// 設定ファイルの読み込み
		if err := viper.ReadInConfig(); err != nil {
			panic(err.Error())
		}

		// 環境ごとの設定ファイルを読み込む
		env := os.Getenv("GO_ENV")
		if env != "" {
			viper.SetConfigName(fmt.Sprintf("config.%s", env))
			if err := viper.MergeInConfig(); err != nil {
				panic(err.Error())
			}
		}

		// configにマッピング
		if err := viper.Unmarshal(&config); err != nil {
			panic(err.Error())
		}
	})
}
