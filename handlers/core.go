package handlers

import (
	"context"
	"log"

	argov1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/rest"
)

func PingHandler(c *gin.Context) {
	c.String(200, "pong")
}

func KubeTenantsGet(c *gin.Context) {
	client, ok := c.MustGet("k8sConn").(rest.RESTClient)
	if !ok {
		log.Panic(ok)
	}

	log.Println("Getting the CRDs")
	// Get and return the YAML from the API server
	result := argov1.ApplicationList{}
	err := client.Get().
		Namespace("default").
		Resource("Applications").
		Do(context.TODO()).
		Into(&result)
	if err != nil {
		panic(err)
	}

	c.YAML(200, result)

}
