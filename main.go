package main

import (
	"log"
	"tatras/clients"
	h "tatras/handlers"

	"github.com/gin-gonic/gin"
)

// Setup
func setupRouter() *gin.Engine {
	r := gin.Default()

	log.Println("Getting the k8s client")
	client, err := clients.GetClientSet()
	if err != nil {
		log.Panic(err)
	}
	r.Use(clients.K8sMiddleware(*client))

	r.GET("/ping", h.PingHandler)
	r.GET("/tenants", h.KubeTenantsGet)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
