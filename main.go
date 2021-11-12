package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantoms158/gin-bookstore/controllers"
	"github.com/phantoms158/gin-bookstore/databases"
	"github.com/phantoms158/gin-bookstore/middlewares"
)

func main() {
	r := gin.Default()

	databases.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})

	r.POST("/login", controllers.Login)


	auth := r.Group("/auth")
	
	auth.Use(middlewares.AuthorizeJWT()) 
	{
		auth.GET("/books", controllers.FindBooks)
		auth.POST("/books", controllers.CreateBook)
		auth.GET("/books/:id", controllers.FindBook)
		auth.PATCH("/books/:id", controllers.UpdateBook)
		auth.DELETE("/books/:id", controllers.DeleteBook)
	}

	
	r.Run()
}