package controller

import "github.com/gin-gonic/gin"

func GetWerSocket(c *gin.Context) {

	c.JSON(200,gin.H{
		"message":"ok",
	})

}