package main

import (
	"main/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	dockerImages := router.Group("/docker-images")
	{
		dockerImages.GET("/local", controllers.GetLocalDockerImages)
		dockerImages.GET("/local/scan", controllers.ScanLocalDockerImage)
		dockerImages.GET("/remote/scan", controllers.ScanRemoteDockerImage)
	}

	fileSystem := router.Group("/file-system")
	{
		fileSystem.GET("/check", controllers.CheckLocalPath)
		fileSystem.GET("/scan", controllers.ScanLocalPath)
	}

	configFiles := router.Group("config-files")
	{
		configFiles.GET("/scan", controllers.ScanLocalConfig)
	}

	repoGit := router.Group("repo-git")
	{
		repoGit.GET("/scan", controllers.ScanRemoteRepoGit)
	}

	router.Run(":4040")
}
