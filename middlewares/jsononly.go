package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JSONOnly middleware that only accept the json format body
func JSONOnly(c *gin.Context) {
	contentType := c.Request.Header["Content-Type"][0]
	if strings.ToLower(contentType) == "application/json" {
		c.Next()
		return
	}

	c.AbortWithError(http.StatusBadRequest, errors.New("add application/json in header"))
}
