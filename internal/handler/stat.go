package handler

import (
	"net/http"
	"shortener"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStat(c *gin.Context) {
	from, err := time.Parse("2006-01-02", c.Query("from"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	to, err := time.Parse("2006-01-02", c.Query("to"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	by := c.Query("by")
	if by != shortener.GroupByDay && by != shortener.GroupByMonth {
		newErrorResponse(c, http.StatusBadRequest, "invalid by param")
		return
	}

	stats, err := h.service.GetStats(by, from, to)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{"stats": stats})
}
