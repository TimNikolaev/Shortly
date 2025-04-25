package handler

import (
	"net/http"
	"shortly"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getStat(c *gin.Context) {
	linkID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

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
	if by != shortly.GroupByDay && by != shortly.GroupByMonth {
		newErrorResponse(c, http.StatusBadRequest, "invalid by param")
		return
	}

	stats, err := h.service.GetStats(uint(linkID), by, from, to)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stats)
}
