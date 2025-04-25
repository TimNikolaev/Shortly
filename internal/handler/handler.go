package handler

import (
	"shortly/internal/service"

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

	link := router.Group("/link", h.userIdentity)
	{
		link.POST("/", h.createLink)
		link.GET("/", h.getAllLinks)
		link.PUT("/:id", h.updateLink)
		link.DELETE("/:id", h.deleteLink)
	}
	router.GET("/:hash", h.goToLink)

	stat := router.Group("/stat", h.userIdentity)
	{
		stat.GET("/:id", h.getStat)
	}

	return router
}
