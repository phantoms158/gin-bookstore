package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantoms158/gin-bookstore/controllers"
	"github.com/phantoms158/gin-bookstore/databases"
)

func main() {
	r := gin.Default()

	databases.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}