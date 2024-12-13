package main

import (
	"net/http"
	"os"

	"github.com/Ayush330/server/config"
	"github.com/Ayush330/server/db"
	"github.com/Ayush330/server/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// create sql pool
	db.InitalizeSql()

	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.DebugLevel)
	gin.DefaultWriter = logrus.StandardLogger().Out
	router := gin.New()
	router.Use(gin.LoggerWithWriter(logrus.StandardLogger().Out))
	router.Use(gin.Recovery())
	router.POST("/createUser", routes.CreateUserHandler)
	router.GET("/getUserData", routes.GetUserDataHandler)
	router.POST("/createGroup", routes.CreateGroupHandler)
	router.POST("/addToGroup", routes.AddToGroupHandler)
	router.POST("/addExpense", routes.AddExpenseHandler)
	router.POST("/getNetExpenseDetailsForAUserAndAGroup", routes.GetNetExpenseForAUserAndGroupHandler)
	http.ListenAndServe(config.GetAddress(), router)

}
