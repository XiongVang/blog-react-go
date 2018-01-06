package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostsGroup(router *gin.Engine) {
	group := router.Group("/posts")

	group.GET("/", func(c *gin.Context) {
		posts, err := GetAllPosts()
		if err != nil {
			log.Println(err)
			c.Status(500)
		} else {
			c.JSON(http.StatusOK, posts)
		}
	})

	group.GET("/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		post, err := GetPost(id)
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, post)
	})

	group.POST("/", func(c *gin.Context) {
		// Read body
		body, err := ioutil.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		if err != nil {
			log.Println("IO ERROR", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// Unmarshal body
		var post Post
		err = json.Unmarshal(body, &post)
		if err != nil {
			log.Println("JSON UNMARSHAL:", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// Create post
		err = post.Create()
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}

	})

	group.PUT("/:id", func(c *gin.Context) {
		// Read body
		body, err := ioutil.ReadAll(c.Request.Body)
		defer c.Request.Body.Close()
		if err != nil {
			log.Println("IO ERROR", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// Unmarshal body
		var post Post
		err = json.Unmarshal(body, &post)
		if err != nil {
			log.Println("JSON UNMARSHAL:", err)
			c.Status(http.StatusInternalServerError)
			return
		}

		// update post
		err = post.Update()
		if err != nil {
			log.Println(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Status(http.StatusOK)
	})
}
