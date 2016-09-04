package main

import (
	"calc-feed-nums/app"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", app.Handlers()))
}
