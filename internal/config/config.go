package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTP ServerConfig `yaml:"http"`
	DB   DBConfig     `yaml:"db"`
	Token TokenConfig `yaml:token`
}

type TokenConfig struct {
	SecretKey  string        `env:"TOKEN_SECRET_KEY" env-required:"true"`
	TimeToLive time.Duration `yaml:"time_to_live"`
}

type ServerConfig struct {
	Port            string        `yaml:"port"`
	Timeout         time.Duration `yaml:"timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
}

type DBConfig struct {
	//Host     string `yaml:"host"`
	//Port     string `yaml:"port"`
	DBName string `yaml:"db_name"`
	//Username string `yaml:"username"`
	//Password string `env:"DB_PASSWORD"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
