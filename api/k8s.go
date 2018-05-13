package api

import (
	"net/http"

	"github.com/mvisonneau/marefo/k8s"

	"github.com/gin-gonic/gin"
)

func GetK8SRunningImages(c *gin.Context) {
	c.JSON(http.StatusOK, k8s.FetchImagesFromRunningPods())
}
