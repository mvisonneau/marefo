package api

import (
  "fmt"

  "github.com/mvisonneau/marefo/config"

  "github.com/gin-gonic/gin"
  "github.com/urfave/cli"
)

func Run(c *cli.Context) {
  switch config.Get().Log.Level {
	case "debug":
		gin.SetMode(gin.DebugMode)
	default:
    fmt.Println("-> Started!")
		gin.SetMode(gin.ReleaseMode)
	}

  r := gin.Default()

  v1 := r.Group("/api/v1/k8s")
  {
    v1.GET("/images", getK8SRunningImages)
  }

  v1 = r.Group("/api/v1/clair")
  {
    v1.GET("/images", getClairKnownImages)
    v1.GET("/images/*image", getClairImageInfo)
  }

  // TODO: Make this configurable more properly
  r.RunTLS(":8443", "./tls/server.crt", "./tls/server.key")
}
