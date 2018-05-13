package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClairKnownImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"images": []string{}})
}

func GetClairImageInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"image": []string{}})
}
