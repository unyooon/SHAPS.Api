package main

import "shaps.api/infrastructure"

func main() {
	d := infrastructure.NewDb()
	r := NewRouting(d)
	r.Run()
}
