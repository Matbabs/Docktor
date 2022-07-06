package controllers

import (
	"main/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLocalDockerImages(c *gin.Context) {
	dockerImages := services.GetLocalDockerImages()
	if len(dockerImages) > 0 && dockerImages != nil {
		c.JSON(http.StatusOK, gin.H{
			"dockerImages": dockerImages,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "no docker local images founded",
		})
	}
}

func ScanLocalDockerImage(c *gin.Context) {
	image := c.Query("image")
	scan := services.ScanLocalDockerImage(image)
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

func ScanRemoteDockerImage(c *gin.Context) {
	image := c.Query("image")
	scan := services.ScanRemoteDockerImage(image)
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
