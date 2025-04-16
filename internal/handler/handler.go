package handler

import (
	"shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	*service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		link := api.Group("/link")
		{
			link.POST("/", h.createLink)
			link.GET("/:id", h.getLink)
			link.GET("/", h.getAllLinks)
			link.PUT("/:id", h.updateLink)
			link.DELETE("/:id", h.deleteLink)
		}
	}
	return router
}
