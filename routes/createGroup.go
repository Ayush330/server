package routes

import (
	"encoding/json"
	"io"

	"github.com/Ayush330/server/db"
	"github.com/Ayush330/server/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateGroupHandler(c *gin.Context) {
	logrus.Warn("In create Group Handler")
	body, _ := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	var requestBody models.CreateGroupPayload
	json.Unmarshal(body, &requestBody)
	IsUserCreated := db.CreateNewGroup(requestBody)
	if IsUserCreated {
		c.JSON(200, gin.H{
			"message": "Group Created",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "failed to create group",
		})
	}
}
