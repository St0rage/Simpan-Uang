package config

import (
	"os"

	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	ApiPort string
	ApiHost string
}

type TokenConfig struct {
	ApplicationName  string
	JwtSignatureKey  string
	JwtSigningMethod *jwt.SigningMethodHMAC
}

type MailConfig struct {
	MailHost     string
	MailPort     string
	MailSender   string
	MailUsername string
	MailPassword string
}

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	DbConfig
	TokenConfig
	MailConfig
}

func (c *Config) readConfigFile() Config {
	err := godotenv.Load()
	utils.PanicIfError(err)

	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
		ApiHost: os.Getenv("API_HOST"),
	}
	c.TokenConfig = TokenConfig{
		ApplicationName:  "SIMPAN-UANG",
		JwtSignatureKey:  "S1MP4N-U4N9",
		JwtSigningMethod: jwt.SigningMethodHS256,
	}
	c.MailConfig = MailConfig{
		MailHost:     os.Getenv("MAIL_HOST"),
		MailPort:     os.Getenv("MAIL_PORT"),
		MailSender:   os.Getenv("MAIL_SENDER"),
		MailUsername: os.Getenv("MAIL_USERNAME"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),
	}

	return *c
}

func NewConfig() Config {
	cfg := Config{}
	return cfg.readConfigFile()
}
