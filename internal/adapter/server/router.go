package server

import (
	"github.com/AlikhanIT/hotel-api/internal/adapter/http"
	"github.com/gin-gonic/gin"
)

func NewRouter(h http.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	h.RegisterRoutes(api)

	return r
}
