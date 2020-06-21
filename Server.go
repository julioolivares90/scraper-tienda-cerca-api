package main

import (
	"os"

	"github.com/julioolivares90/tiendaCercaApi/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	runServer()
}

func runServer() {

	server := echo.New()

	server.GET("api/find-lugar", handlers.GetLugar)

	server.Logger.Fatal(server.Start(":" + getPort()))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	return port
}
