package api

import (
	"net/http"

	"github.com/Satssuki/Go-Service-Boilerplate/helpers"
)

// JSONResponse function supposed to be transform fetched struct to bytes and respond it
func JSONResponse(statusCode int, writer http.ResponseWriter, data interface{}) error {
	writer.WriteHeader(statusCode)
	writer.Header().Add("Content-Type", "application/json")
	return helpers.JSONNewEncoder(writer).Encode(data)
}
