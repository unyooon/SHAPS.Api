package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// 環境変数読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := InitializeHandler(os.Getenv("DSN"))
	r.Run()
}
