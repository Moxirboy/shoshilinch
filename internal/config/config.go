package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload" 
)



var instance Config
func Load() *Config{
	if err:=env.Parse(&instance);err!=nil{
		log.Fatalf("could not load env file error:"+err.Error())
		return nil
	}
	return &instance
}

type Config struct{
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`
	Logger Logger
	Server Server
	JWT	   JWT
}
type Postgres struct{
	Host 	 string `env:"POSTGRES_HOST"`
	Port 	 string `env:"POSTGRES_PORT"`
	Database string `env:"POSTGRES_DATABASE"`
	Username string `env:"POSTGRES_USERNAME"`
	Password string `env:"POSTGRES_PASSWORD"`
}
type (
	
	Server struct {
	Environment       string `env:"SERVER_ENVIRONMENT"`
	Port              uint16 `env:"ADMIN_PORT"`
	MaxConnectionIdle uint16 `env:"SERVER_MAX_CONNECTION_IDLE"`
	Timeout           uint16 `env:"SERVER_TIMEOUT"`
	Time              uint16 `env:"SERVER_TIME"`
	MaxConnectionAge  uint16 `env:"SERVER_MAX_CONNECTION_AGE"`
}
	Logger struct {
		Level string `env:"LOGGER_LEVEL"`
		Encoding string `env:"LOGGER_ENCODING"`
	}
	 JWT struct{
		SecretKey string `env:"SECRET_KEY"`
		SecretKeyExpireMinutes   uint16 `env:"JWT_ADMIN_SECRET_KEY_EXPIRE_MINUTES"`
	}

)