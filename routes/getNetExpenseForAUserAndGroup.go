package routes

import (
	"encoding/json"
	"io"

	"github.com/Ayush330/server/db"
	"github.com/Ayush330/server/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetNetExpenseForAUserAndGroupHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	var requestBody models.ExpenseDetailsUserGroupPayload
	json.Unmarshal(body, &requestBody)
	Data, err := db.GetNetExpenseDetailsForAUserForAGroup(requestBody)
	if err != nil {
		logrus.Error("Error in GetNetExpenseForAUserAndGroupHandler: ", err.Error())
		c.JSON(200, gin.H{
			"message": "failed to fetch data",
		})
	} else {
		c.JSON(200, gin.H{
			"message": Data,
		})
	}
}
