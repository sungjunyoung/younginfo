package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func main() {
	client := resty.New()

	version := os.Getenv("VERSION")
	if version == "" {
		version = "unknown"
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		ratingsVersion, err := client.R().Get("http://ratings:9080")
		if err != nil {
			return c.String(http.StatusInternalServerError, "reviews unavailable")
		}

		return c.String(http.StatusOK,
			fmt.Sprintf("reviews: %s with ratings: %s", version, ratingsVersion))
	})

	e.Logger.Fatal(e.Start(":9080"))
}
