package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown"
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ratings version: "+version)
	})

	e.Logger.Fatal(e.Start(":9080"))
}
