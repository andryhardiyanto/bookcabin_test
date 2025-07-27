package config

import (
	"time"

	"github.com/spf13/viper"
)

type config struct {
	Env               string        `mapstructure:"ENV"`
	AppName           string        `mapstructure:"APP_NAME"`
	AppPort           string        `mapstructure:"APP_PORT"`
	DBPath            string        `mapstructure:"DB_PATH"`
	DBMaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBConnMaxIdleTime time.Duration `mapstructure:"DB_CONN_MAX_IDLE_TIME"`
	DBConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

var Cfg *config

func InitConfig() error {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return err
	}

	return nil
}
