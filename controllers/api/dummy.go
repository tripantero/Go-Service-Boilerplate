package v1

import (
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	"github.com/Satssuki/Go-Service-Boilerplate/seeders"
	"github.com/gin-gonic/gin"
)

// Dummy no context
func Dummy(c *gin.Context) {
	api.JSONResponse(200, c.Writer, seeders.Starter())
}
