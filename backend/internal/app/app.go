package app

import (
	"log/slog"

	httpapp "github.com/wolkodaf/todo/backend/internal/app/http"
	"github.com/wolkodaf/todo/backend/internal/config"
	"github.com/wolkodaf/todo/backend/internal/services"
	"github.com/wolkodaf/todo/backend/internal/transport"
	todoHandler "github.com/wolkodaf/todo/backend/internal/transport/handlers/todos"
)

type App struct {
	HTTPServer *httpapp.App
	log        *slog.Logger
}

func New(log *slog.Logger, cfg *config.Config) *App {
	app := &App{log: log}

	// В api_gateway тут инициализировались gRPC-клиенты к микросервисам.
	// Здесь пропускаем этот шаг и используем простой in-memory сервис.
	todoService := services.NewTodoService()

	router := transport.NewRouter(cfg)
	apiGroup := router.Group("/api")

	todoHndlr := todoHandler.New(log, todoService)
	todoHndlr.RegisterRoutes(apiGroup)

	app.HTTPServer = httpapp.New(log, cfg.HTTPServer.Port, router)

	return app
}
