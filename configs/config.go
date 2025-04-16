package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig
	AuthConfig
}

type DBConfig struct {
	DSN string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		DBConfig: DBConfig{
			DSN: os.Getenv("DSN"),
		},
		AuthConfig: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}, nil
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
