package domain

// Todo — модель задачи.
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// CreateTodoRequest — тело запроса на создание задачи.
type CreateTodoRequest struct {
	Title string `json:"title" binding:"required"`
}

// UpdateTodoRequest — тело запроса на обновление задачи.
type UpdateTodoRequest struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}
