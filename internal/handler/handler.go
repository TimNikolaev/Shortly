package handler

import (
	"shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		link := api.Group("/link")
		{
			link.POST("/", h.createLink)
			link.GET("/:hash", h.goToLink)
			link.GET("/", h.getAllLinks)
			link.PUT("/:id", h.updateLink)
			link.DELETE("/:id", h.deleteLink)
		}
	}
	return router
}
