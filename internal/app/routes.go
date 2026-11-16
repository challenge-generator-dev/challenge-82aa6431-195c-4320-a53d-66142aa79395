package app

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	// Sample route - will be replaced by domain code
	router.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sample endpoint",
		})
	})
}