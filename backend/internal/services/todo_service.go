package services

import (
	"errors"
	"sync"

	"github.com/wolkodaf/todo/backend/internal/domain"
)

var ErrNotFound = errors.New("todo not found")

// TodoService — простое in-memory хранилище задач.
// В api_gateway на этом месте были бы вызовы gRPC-микросервисов.
type TodoService struct {
	mu     sync.RWMutex
	nextID int
	items  map[int]domain.Todo
}

func NewTodoService() *TodoService {
	s := &TodoService{
		nextID: 1,
		items:  make(map[int]domain.Todo),
	}
	// Немного стартовых данных, чтобы фронт был не пустым.
	s.Create("Первая задача")
	s.Create("Вторая задача")
	return s
}

func (s *TodoService) List() []domain.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]domain.Todo, 0, len(s.items))
	for _, t := range s.items {
		result = append(result, t)
	}
	// Стабильный порядок по ID.
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[j].ID < result[i].ID {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}

func (s *TodoService) Create(title string) domain.Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := domain.Todo{ID: s.nextID, Title: title, Done: false}
	s.items[todo.ID] = todo
	s.nextID++
	return todo
}

func (s *TodoService) Update(id int, title *string, done *bool) (domain.Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.items[id]
	if !ok {
		return domain.Todo{}, ErrNotFound
	}
	if title != nil {
		todo.Title = *title
	}
	if done != nil {
		todo.Done = *done
	}
	s.items[id] = todo
	return todo, nil
}

func (s *TodoService) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.items[id]; !ok {
		return ErrNotFound
	}
	delete(s.items, id)
	return nil
}
