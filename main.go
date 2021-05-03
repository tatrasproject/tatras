package main

import (
	h "tatras/handlers"

	"github.com/gin-gonic/gin"
)

// Setup
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", h.PingHandler)
	r.GET("/tenants", h.KubeTenantsGet)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
