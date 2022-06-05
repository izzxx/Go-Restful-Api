package config

import "os"

var (
	JwtSecretkey string
	DbPort       string
	ServerPort   string
)

func InitConfig() {
	JwtSecretkey = os.Getenv("JWT_KEY")
	DbPort = os.Getenv("DB_URL")
	ServerPort = os.Getenv("SERVER_PORT")
}
