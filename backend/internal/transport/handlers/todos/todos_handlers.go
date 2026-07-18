package todos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wolkodaf/todo/backend/internal/domain"
	"github.com/wolkodaf/todo/backend/internal/lib/response"
	"github.com/wolkodaf/todo/backend/internal/services"
)

func (h *Handler) list(c *gin.Context) {
	response.Success(c, h.service.List())
}

func (h *Handler) create(c *gin.Context) {
	var input domain.CreateTodoRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, err)
		return
	}

	todo := h.service.Create(input.Title)
	response.Success(c, todo)
}

func (h *Handler) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var input domain.UpdateTodoRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, err)
		return
	}

	todo, err := h.service.Update(id, input.Title, input.Done)
	if err != nil {
		response.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	response.Success(c, todo)
}

func (h *Handler) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, "invalid id parameter")
		return
	}

	if err := h.service.Delete(id); err != nil {
		if err == services.ErrNotFound {
			response.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		response.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(c, gin.H{"deleted": id})
}
