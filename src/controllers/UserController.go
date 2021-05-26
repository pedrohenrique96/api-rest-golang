package controllers

import (
	"crud/src/database"
	"crud/src/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowAll(c *gin.Context) {
	db := database.GetDatabase()

	var u []models.User

	err := db.Find(&u).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot find users: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, u)
}

func Save(c *gin.Context) {
	db := database.GetDatabase()

	var u models.User

	err := c.ShouldBindJSON(&u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&u).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, u)
}
