package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	const port = ":5555"

	router := gin.Default()
	helloGroup(router)
	router.Static("/static", "./public/static")
	router.StaticFile("/", "./public/index.html")
	log.Fatal(router.Run(port))
}

func helloGroup(router *gin.Engine) {
	hello := router.Group("/hi")

	hello.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	hello.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}
