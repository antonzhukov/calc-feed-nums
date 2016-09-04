package app

import (
	"calc-feed-nums/app/service"
	"net/http"
)

func Handlers() *http.ServeMux {
	// Register application routes and corresponding services
	routes := http.NewServeMux()
	routes.HandleFunc("/numbers", service.CalculateNumbersFeed)

	return routes
}
