package app

import (
	"calc-feed-nums/app/service"
	"net/http"
)

func RegisterRoutes() {
	// Register application routes and corresponding services
	http.HandleFunc("/numbers", service.CalculateNumbersFeed)
}
