package routes

import (
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user Created",
	})
}
