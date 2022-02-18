package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "shaps.api/docs"
	"shaps.api/domain/setting"
)

// @title SHAPS API
// @version version(1.0)
// @description SHAPS API
// @license.name yuta

// @host localhost:3000
// @BasePath /
func main() {
	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s := setting.Setting{
		Dsn:          os.Getenv("DSN"),
		Port:         os.Getenv("PORT"),
		B2CJwkUri:    os.Getenv("B2C_JWK_ENDPOINT"),
		B2CTenantId:  os.Getenv("B2C_TENANT_ID"),
		B2CClientId:  os.Getenv("B2C_CLIENT_ID"),
		StripeApiKey: os.Getenv("STRIPE_API_KEY"),
	}
	r := InitializeHandler(s)
	r.Run()
}
