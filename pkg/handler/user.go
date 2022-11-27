package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUsername(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	username, err := h.services.User.Get(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, username)
}
