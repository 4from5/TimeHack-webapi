package handler

import (
	"fmt"
	webapi "github.com/4from5/TimeHack-webapi"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getCategories(c *gin.Context) {

}

func (h *Handler) createCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input webapi.Category
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.createCategory:", input)
	id, err := h.services.Categories.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"category_id": id,
	})
	//	input.UserId, _ = strconv.Atoi(cast.ToString(id)) //это шок просто ужас

}
