package main

import (
	"net/http"

	"github.com/Ayush330/server/config"
	"github.com/Ayush330/server/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/createUser", routes.CreateUserHandler)
	router.GET("/getUserData", routes.GetUserDataHandler)

	http.ListenAndServe(config.GetAddress(), router)
}
