package api

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
	"github.com/gin-gonic/gin"
)

// JSONResponse function supposed to be transform fetched struct to bytes and respond it
func JSONResponse(statusCode int, writer http.ResponseWriter, data *gin.H) error {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")
	return helpers.JSONNewEncoder(writer).Encode(*data)
}
