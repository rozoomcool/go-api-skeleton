package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Port       int           `mapstructure:"PORT"`
	DBHost     string        `mapstructure:"DB_HOST"`
	DBPort     int           `mapstructure:"DB_PORT"`
	DBPassword string        `mapstructure:"DB_PASSWORD"`
	DBName     string        `mapstructure:"DB_NAME"`
	JWTSecret  string        `mapstructure:"JWT_SECRET"`
	AccessTTL  time.Duration `mapstructure:"ACCESS_TTL"`
	RefreshTTL time.Duration `mapstructure:"REFRESH_TTL"`
	WebhookURL string        `mapstructure:"WEBHOOK_URL"`
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.AutomaticEnv()
	_ = v.ReadInConfig()

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
