package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScanLocalConfig(c *gin.Context) {
	path := c.Query("path")
	scan := services.ScanLocalConfig(path)
	if scan != nil {
		c.JSON(http.StatusOK, gin.H{
			"scan": scan,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "scan aborted",
		})
	}
}
