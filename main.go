package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"shaps.api/domain/setting"
)

func main() {

	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := InitializeHandler(setting.Setting{
		Port: os.Getenv("PORT"),
		Dsn:  os.Getenv("DSN"),
	})
	r.Run()
}
