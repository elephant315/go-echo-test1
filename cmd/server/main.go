package main

import (
	_ "github.com/elephant315/go-echo-test1/docs"
	"github.com/elephant315/go-echo-test1/internal/adapters/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Flight Itinerary API
// @version 1.0
// @description This is a simple API to reconstruct flight itineraries from given tickets.
// @host localhost:8080
// @BasePath /
func main() {
	e := echo.New()

	// Swagger endpoint
	e.GET("/apidoc/*", echoSwagger.WrapHandler)

	// Start server with HTTP handlers
	http.RegisterHandlers(e)
	e.Logger.Fatal(e.Start(":8080"))
}
