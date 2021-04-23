package main

import (
	clients "tatras/clients"
	h "tatras/handlers"

	"github.com/gin-gonic/gin"
)

// Setup
func setupRouter() *gin.Engine {
	r := gin.Default()

	//TODO conditionally create the in-cluster version of the client
	clients.GetClientSetFromStandalone()

	r.GET("/ping", h.PingHandler)
	
	r.GET("/argo/:application", h.GetArgoApplication)
	r.POST("/argo/:application", h.CreateArgoApplication)

	return r
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8080")
}
