package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScanRemoteRepoGit(c *gin.Context) {
	uri := c.Query("uri")
	scan := services.ScanRemoteRepoGit(uri)
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
