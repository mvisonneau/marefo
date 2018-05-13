package api

import (
	"net/http"

  "github.com/mvisonneau/marefo/clair"
  "github.com/mvisonneau/marefo/config"

	"github.com/gin-gonic/gin"
)

func getClairKnownImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"images": []string{}})
}

func getClairImageInfo(c *gin.Context) {
  cl, err := clair.NewClient(config.Get().Clair.Endpoint)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"clair_get_client_error": err })
	}

  vulnerabilities, err := cl.Analyze(trimLeftChar(c.Param("image")))
	if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"clair_analyze_error": err })
	} else {
    c.JSON(http.StatusOK, gin.H{"clair_info": vulnerabilities})
  }
}

func trimLeftChar(s string) string {
    for i := range s {
        if i > 0 {
            return s[i:]
        }
    }
    return s[:0]
}
