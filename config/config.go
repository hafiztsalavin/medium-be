package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	Port           uint16 `env:"PORT,default=8888"`
	Env            string `env:"ENV"`
	DatabaseConfig DatabaseOption
}

type DatabaseOption struct {
	Username string `env:"DATABASE_USERNAME,required"`
	Password string `env:"DATABASE_PASSWORD,required"`
	Host     string `env:"DATABASE_HOST,default=localhost"`
	Port     string `env:"DATABASE_PORT,default=5432"`
	Name     string `env:"DATABASE_NAME,required"`
}

func NewConfig() *Config {
	var cfg Config
	gotenv.Load(".env")

	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
