package config

import (
	"github.com/Audibene-GMBH/ta.go-hexagonal-skeletor/pkg/env"
)

type Config struct {
	App   AppConfig
	Db    DbConfig
	Redis RedisConfig
}

type AppConfig struct {
	Name string `env:"APP_NAME" default:"ta.go-hexagonal-skeletor"`
	Port string `env:"APP_PORT" default:"3000"`
	Env  string `env:"APP_ENV" default:"development"`
	Host string `env:"APP_HOST" default:"localhost"`
}

type DbConfig struct {
	Name     string `env:"RDS_DBNAME"`
	Host     string `env:"RDS_HOST"`
	Port     string `env:"RDS_PORT"`
	Username string `env:"RDS_USERNAME"`
	Password string `env:"RDS_PASSWORD"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
}

func MustLoad(filenames ...string) *Config {
	conf := Config{}
	env.Load(filenames...)
	err := env.Marshal(&conf)
	if err != nil {
		panic(err)
	}
	return &conf
}
