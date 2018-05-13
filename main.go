package main

import (
	"github.com/mvisonneau/marefo/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1/k8s")
	{
		v1.GET("/images", api.GetK8SRunningImages)
	}

	v1 = r.Group("/api/v1/clair")
	{
		v1.GET("/images", api.GetClairKnownImages)
		v1.GET("/images/:name", api.GetClairImageInfo)
	}

	// TODO: Make this configurable more properly
	r.RunTLS(":8443", "./tls/server.crt", "./tls/server.key")
}
