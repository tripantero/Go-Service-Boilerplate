package route

import (
	v1 "github.com/Satssuki/Go-Service-Boilerplate/controllers/api/v1"
	"github.com/Satssuki/Go-Service-Boilerplate/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRouter Register each route here
func SetupRouter() *gin.Engine {
	router := gin.Default()	
	router.Use(middlewares.AppliAllCORS)
	
	v1g := router.Group("/api/v1")
	{
		v1g.POST("/user", v1.InsertUser)
	}
	return router
}
