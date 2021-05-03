package handlers

import (
	"context"
	"tatras/clients"

	argov1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.String(200, "pong")
}

func KubeTenantsGet(c *gin.Context) {
	client, err := clients.GetClientSet()
	if err != nil {
		panic(err)
	}

	// Get and return the YAML from the API server
	result := argov1.ApplicationList{}
	err = client.Get().
		Namespace("default").
		Resource("Applications").
		Do(context.TODO()).
		Into(&result)
	if err != nil {
		panic(err)
	}

	c.YAML(200, result)

}
