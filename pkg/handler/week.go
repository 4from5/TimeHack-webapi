package handler

import (
	webApi "github.com/4from5/TimeHack-webapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getWeek(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input webApi.WeekRequest
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	week, err := h.services.Week.GetDays(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, webApi.Week{week})
}
