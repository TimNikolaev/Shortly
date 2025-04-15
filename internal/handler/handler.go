package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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
