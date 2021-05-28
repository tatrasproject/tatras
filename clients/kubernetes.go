package clients

import (
	"fmt"

	argov1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/gin-gonic/gin"
	discovery "github.com/gkarthiks/k8s-discovery"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func GetClientSet() (*rest.RESTClient, error) {
	k8s, _ := discovery.NewK8s()
	namespace, _ := k8s.GetNamespace()
	version, _ := k8s.GetVersion()
	fmt.Printf("Specified Namespace: %s\n", namespace)
	fmt.Printf("Version of running Kubernetes: %s\n", version)

	config := k8s.RestConfig

	argov1.AddToScheme(scheme.Scheme)
	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: argov1.SchemeGroupVersion.Group, Version: argov1.SchemeGroupVersion.Version}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	// return the clientset
	return rest.UnversionedRESTClientFor(&crdConfig)

}

// ApiMiddleware will add the k8s connection to the context
// H/T https://stackoverflow.com/a/34053075
func K8sMiddleware(k8s rest.RESTClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("k8sConn", k8s)
		c.Next()
	}
}
