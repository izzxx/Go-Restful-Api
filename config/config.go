package config

import "os"

var (
	JwtSecretkey string
	ServerPort   string
	DbHost       string
	DbPort       string
	DbName       string
	DbPassword   string
	DbUser       string
)

func InitConfig() {
	JwtSecretkey = os.Getenv("JWT_KEY")
	ServerPort = os.Getenv("SERVER_PORT")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbUser = os.Getenv("DB_USER")
}
