package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckLocalPath(c *gin.Context) {
	path := c.Query("path")
	stdout := services.CheckLocalPath(path)
	if len(stdout) > 0 && stdout != nil {
		c.JSON(http.StatusOK, gin.H{
			"stdout": stdout,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "command error",
		})
	}
}

func ScanLocalPath(c *gin.Context) {
	path := c.Query("path")
	scan := services.ScanLocalPath(path)
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
