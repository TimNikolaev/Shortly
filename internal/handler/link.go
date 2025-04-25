package handler

import (
	"net/http"
	"shortly"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

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

	c.JSON(http.StatusOK, link)
}

func (h *Handler) goToLink(c *gin.Context) {
	hashLink := c.Param("hash")

	link, err := h.service.GoToLink(hashLink)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, link.URL)
}

type getAllLinksResponse struct {
	Links []shortly.Link `json:"links"`
	Count int64          `json:"count"`
}

func (h *Handler) getAllLinks(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	links, count, err := h.service.GetAllLinks(userID, limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLinksResponse{Links: links, Count: count})
}

type LinkUpdateRequest struct {
	URL  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

func (h *Handler) updateLink(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	linkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input LinkUpdateRequest

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	link, err := h.service.UpdateLink(uint(userID), uint(linkID), input.URL, input.Hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, link)
}

func (h *Handler) deleteLink(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	linkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.DeleteLink(uint(userID), uint(linkID))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
