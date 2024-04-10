package main

import (
	"log"
	"net/http"
	"os"

	"os/exec"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		cmd := exec.Command("whoami")
		output, err := cmd.Output()
		if err != nil {
			 log.Println(err)
		}
		c.HTML(http.StatusOK, "index.tmpl.html", output)
	})

	router.Run(":" + port)
}
