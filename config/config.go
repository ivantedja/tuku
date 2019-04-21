package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

//Config represent global config
type Config struct {
	Database struct {
		Username string `env:"DATABASE_USERNAME,required"`
		Password string `env:"DATABASE_PASSWORD,required"`
		Host     string `env:"DATABASE_HOST,default=localhost"`
		Port     string `env:"DATABASE_PORT,default=3306"`
		Name     string `env:"DATABASE_NAME,required"`
		Charset  string `env:"DATABASE_CHARSET,default=utf8"`
		Pool     int    `env:"DATABASE_POOL,default=50"`
	}
	Slack struct {
		Token     string `env:"SLACK_TOKEN,required"`
		ChannelID string `env:"SLACK_CHANNEL_ID,required"`
	}
}

func NewConfig() (Config, error) {
	gotenv.Load(".env")
	var cfg Config
	err := envdecode.Decode(&cfg)
	return cfg, err
}
