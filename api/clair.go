package api

import (
	"net/http"
  "fmt"
  "encoding/json"

  "github.com/mvisonneau/marefo/clair"
  "github.com/mvisonneau/marefo/config"

  "github.com/coreos/clair/api/v3/clairpb"
	"github.com/gin-gonic/gin"
  "k8s.io/api/admission/v1beta1"
  "k8s.io/api/core/v1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getClairKnownImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"images": []string{}})
}

func getClairImageInfo(c *gin.Context) {
  vulns, err := clairAnalyzeImage(trimLeftChar(c.Param("image")))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"clair_error": err })
	} else {
    c.JSON(http.StatusOK, gin.H{"clair_data": vulns})
  }
}

func postClairAdmitImage(c *gin.Context) {
  reviewResponse := &v1beta1.AdmissionResponse{
    Allowed: true,
  }
  exitCode := http.StatusOK
  req := v1beta1.AdmissionReview{}
  if err := c.ShouldBindJSON(&req); err == nil {
		reviewResponse = &v1beta1.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
    exitCode = http.StatusBadRequest
	} else {
    pod := v1.Pod{}
    if err := json.Unmarshal(req.Request.Object.Raw, &pod); err != nil {
      reviewResponse = &v1beta1.AdmissionResponse{
  			Result: &metav1.Status{
  				Message: err.Error(),
  			},
  		}
      exitCode = http.StatusBadRequest
  	}
    for _, container := range pod.Spec.Containers {
      vulns, err := clairAnalyzeImage(container.Image)
      if err != nil || len(vulns) > 0 {
        reviewResponse = &v1beta1.AdmissionResponse{
          Result: &metav1.Status{
            Message: fmt.Sprintf("Image %s is vulnerable", container.Image),
          },
        }
        exitCode = http.StatusBadRequest
        break
      }
    }
	}

  c.JSON(exitCode, reviewResponse)
}

func clairAnalyzeImage(image string) ([]*clairpb.Vulnerability, error) {
  cl, err := clair.NewClient(config.Get().Clair.Endpoint)
  if err != nil {
    return nil, err
	}

  vulns, err := cl.Analyze(image)
	if err != nil {
    return nil, err
	}

  return vulns, nil
}

func trimLeftChar(s string) string {
    for i := range s {
        if i > 0 {
            return s[i:]
        }
    }
    return s[:0]
}
