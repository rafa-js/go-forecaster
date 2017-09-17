package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"./views"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		//log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("api/forecasts/pending", views.GetPendingForecasts)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "SHU", nil)
	})

	router.Run(":" + port)
}
