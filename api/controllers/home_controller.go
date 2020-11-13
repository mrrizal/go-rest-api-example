package controllers

import (
	"net/http"

	"github.com/mrrizal/go-rest-api-example/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome this awesome API")
}
