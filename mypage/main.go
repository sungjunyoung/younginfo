package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	client := resty.New()

	e := echo.New()
	e.GET("/mypage", func(c echo.Context) error {
		reviewsVersion, err := client.R().Get("http://reviews:9080")
		if err != nil {
			return c.String(http.StatusInternalServerError, "reviews unavailable")
		}

		detailsVersion, err := client.R().Get("http://details:9080")
		if err != nil {
			return c.String(http.StatusInternalServerError, "details unavailable")
		}

		return c.String(http.StatusOK,
			"REVIEWS: "+reviewsVersion.String()+"\n"+
				"DETAILS: "+detailsVersion.String()+"\n",
		)
	})

	e.Logger.Fatal(e.Start(":9080"))
}
