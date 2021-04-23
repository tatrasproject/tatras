package handlers

import (
	"github.com/gin-gonic/gin"
	argo "tatras/clients/argo"
)

func PingHandler(c *gin.Context) {
	c.String(200, "pong")
}

// Retrieve info about an application
func GetArgoApplication(c *gin.Context) {
	app := c.Param("application")
	response := argo.GetApplication(app)
	c.JSON(200, string(response))
}

// Create an application
func CreateArgoApplication(c *gin.Context) {
	app := c.Param("application")
	response := argo.CreateApplication(app)
	c.JSON(200, string(response))
}
