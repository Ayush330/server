package routes

import (
	"github.com/gin-gonic/gin"
)

func GetUserDataHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "data fetched",
	})
}
