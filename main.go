package main

import (
	"net/http"
	"time"

	"github.com/eduardo-js/go-gin-api/src/router"
)

func main() {
	api := &http.Server{
		Addr:         ":4000",
		Handler:      router.RouteHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	api.ListenAndServe()
}
