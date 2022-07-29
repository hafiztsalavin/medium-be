package config

import (
	"medium-be/internal/constants"
	"os"

	"github.com/joeshaw/envdecode"
	"github.com/subosito/gotenv"
)

type Config struct {
	Port           uint16 `env:"PORT,default=8888"`
	Env            string `env:"ENV"`
	DatabaseConfig DatabaseOption
	RedisConfig    RedisOption
}

type DatabaseOption struct {
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Host     string `env:"DB_HOST,default=localhost"`
	Port     string `env:"DB_PORT,default=5432"`
	Name     string `env:"DB_NAME,required"`
	Timeout  int    `env:"DB_TIMEOUT,required"`
}

type RedisOption struct {
	Addr     string `env:"REDIS_ADDR,required"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB,default=0"`
}

func NewConfig() *Config {
	var cfg Config
	gotenv.Load(".env")

	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}

	constants.JWT_ACCESS_KEY = os.Getenv("JWT_ACCESS_KEY")
	constants.JWT_REFRESH_KEY = os.Getenv("JWT_REFRESH_KEY")

	return &cfg
}
