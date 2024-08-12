package controllers

import (
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePackage(c *gin.Context) {
	var input models.Package
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

func UpdatePackage(c *gin.Context) {
	var existingPackage models.Package
	if err := database.DB.First(&existingPackage, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Package not found!"})
		return
	}
	var input models.Package
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&existingPackage).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": existingPackage})
}

func DeletePackage(c *gin.Context) {
	var existingPackage models.Package
	if err := database.DB.First(&existingPackage, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Package not found!"})
		return
	}
	database.DB.Delete(&existingPackage)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
