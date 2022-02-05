package main

import "os"

func main() {
	r := InitializeHandler(os.Getenv("DSN"))
	r.Run()
}
