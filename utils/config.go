package utils

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	MIGRATIONURL        string        `mapstructure:"MIGRATION_URL"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokkenStructureKey  string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`

	BaseUrl string `mapstructure:"base_url"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
	GINMODE   string `mapstructure:"GIN_MODE"`
}

func LoadConfiguration(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return

}
