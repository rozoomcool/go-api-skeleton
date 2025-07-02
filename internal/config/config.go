package config

import "time"

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
