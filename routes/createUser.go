package routes

import (
	"encoding/json"
	"io"

	"github.com/Ayush330/server/db"
	"github.com/Ayush330/server/models"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	var requestBody models.CreateUserPayload
	json.Unmarshal(body, &requestBody)
	IsUserCreated := db.CreateNewUser(requestBody)
	if IsUserCreated {
		c.JSON(200, gin.H{
			"message": "user Created",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "failed to create user",
		})
	}
}
