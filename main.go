package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"shaps.api/domain/setting"
	"shaps.api/middleware"
)

func main() {
	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s := setting.Setting{
		Dsn:         os.Getenv("DSN"),
		Port:        os.Getenv("PORT"),
		B2CJwkUri:   os.Getenv("B2C_JWK_ENDPOINT"),
		B2CTenantId: os.Getenv("B2C_TENANT_ID"),
		B2CClientId: os.Getenv("B2C_CLIENT_ID"),
	}
	r := InitializeHandler(s)
	r.Run()
}
