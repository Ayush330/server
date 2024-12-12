package routes

import (
	"encoding/json"
	"io"

	"github.com/Ayush330/server/db"
	"github.com/Ayush330/server/models"
	"github.com/gin-gonic/gin"
)

func AddToGroupHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	var requestBody models.AddToGroupPayload
	json.Unmarshal(body, &requestBody)
	IsUserCreated := db.AddToGroup(requestBody)
	if IsUserCreated {
		c.JSON(200, gin.H{
			"message": "Added to group",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "failed to add to  group",
		})
	}
}
