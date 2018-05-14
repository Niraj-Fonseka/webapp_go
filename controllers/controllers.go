package controllers

import (
	"webapp_go/models"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"Health": "Health is Ok !",
	})
}

func CreateCreds(c *gin.Context) {
	models.CreateCredentials("NirajUsername", "Paswword")
}