package main

import (
	"calc-feed-nums/app"
	"log"
	"net/http"
)

func main() {
	app.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
