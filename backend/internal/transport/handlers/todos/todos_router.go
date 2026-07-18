package todos

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/wolkodaf/todo/backend/internal/services"
)

type Handler struct {
	log     *slog.Logger
	service *services.TodoService
}

func New(log *slog.Logger, service *services.TodoService) *Handler {
	return &Handler{
		log:     log,
		service: service,
	}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	todosGroup := rg.Group("/todos")
	{
		todosGroup.GET("", h.list)
		todosGroup.POST("", h.create)
		todosGroup.PATCH("/:id", h.update)
		todosGroup.DELETE("/:id", h.delete)
	}
}
