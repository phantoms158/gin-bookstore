package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phantoms158/gin-bookstore/models"
)

// POST /login
// login
func Login(c *gin.Context) {
	var input models.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// // Login Information
	// user := models.User{Username: input.Username, Password: input.Password}
	// databases.DB.Create(&user)

	// c.JSON(http.StatusOK, gin.H{"data": user})
	token := models.GenerateToken(input.Username, true)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, nil)
	}
}
