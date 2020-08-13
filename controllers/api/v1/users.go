package v1

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/Satssuki/Go-Service-Boilerplate/helpers/api"
	v1s "github.com/Satssuki/Go-Service-Boilerplate/services/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// InsertUser sample controller to perform insert user function
func InsertUser(c *gin.Context) {
	defer c.Request.Body.Close()

	service := v1s.CreateUserService()
	err := helpers.ReadByteAndParse(c.Request.Body, &service.User)

	if err == nil {
		message, err := service.Insert()
		if err == nil {
			api.JSONResponse(http.StatusOK, c.Writer, gin.H{
				"status":  "ok",
				"message": message,
			})
			log.
				Info().
				Msgf("User registered with name: %v", service.User.Name)
			return
		} else {
			log.
				Warn().
				Str("error", err.Error()).
				Msg("InsertUser-> service.Insert failure")
		}
	} else {
		log.
			Warn().
			Str("error", err.Error()).
			Msg("InsertUser-> failed to parse byte")
	}

	api.JSONResponse(http.StatusBadRequest, c.Writer, gin.H{
		"status":  "failure",
		"message": err,
	})
}
