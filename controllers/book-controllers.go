package controllers

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/phantoms158/gin-bookstore/databases"
	"github.com/phantoms158/gin-bookstore/models"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	claims, exists := c.Get("user_data")
	if !exists {
		c.JSON(http.StatusOK, gin.H{"data": exists})
	}
	user_data := claims.(jwt.MapClaims)
	fmt.Printf("Key: %v \n", user_data["name"].(string))
	var books []models.Book
	databases.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input models.CreateBookInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Book
	book := models.Book{Title: input.Title, Author: input.Author}
	databases.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// Get /books/:id
// Find a Book
func FindBook(c *gin.Context) {
	// Get model if exists
	var book models.Book
	err := databases.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	// Get model if exists
	var book models.Book
	err := databases.DB.Where("id = ?", c.Param("id")).First(&book).Error;
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input models.UpdateBookInput
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	databases.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	// Delete model if exists
	var book models.Book
	err := databases.DB.Where("id = ?", c.Param("id")).First(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	databases.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}