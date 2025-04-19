package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createLink(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input LinkCreateRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	link, err := h.service.CreateLink(userID, input.URL)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{"hash": link.Hash})

}

func (h *Handler) getLink(c *gin.Context) {

}

func (h *Handler) getAllLinks(c *gin.Context) {

}

func (h *Handler) updateLink(c *gin.Context) {

}

func (h *Handler) deleteLink(c *gin.Context) {

}
