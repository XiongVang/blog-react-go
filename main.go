package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	const port = ":3000"

	ginRouter := gin.Default()
	ginRouter.Static("/", "./public")
	log.Fatal(ginRouter.Run(port))
}
