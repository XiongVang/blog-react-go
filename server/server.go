package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	const port = ":5555"

	router := gin.Default()
	PostsGroup(router)
	router.Static("/static", "../public/")
	router.StaticFile("/", "../public/index.html")

	log.Fatal(router.Run(port))
}
