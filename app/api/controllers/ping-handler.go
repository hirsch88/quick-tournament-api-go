package controllers

import "github.com/gin-gonic/gin"

func PongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
