package routes

import (
	"encoding/json"
	"io"

	"github.com/Ayush330/server/db"
	"github.com/Ayush330/server/models"
	"github.com/gin-gonic/gin"
)

func AddExpenseHandler(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	var requestBody models.AddExpensePayload
	json.Unmarshal(body, &requestBody)
	IsUserCreated := db.AddExpense(requestBody)
	if IsUserCreated {
		c.JSON(200, gin.H{
			"message": "Added Expense",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Failed To Add Expense",
		})
	}
}
