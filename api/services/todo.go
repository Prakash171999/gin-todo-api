package services

import (
	"boilerplate-api/api/repository"
)

// UserService -> struct
type TodoService struct {
	repository repository.TodoRepository
}

// NewUserService -> creates a new Userservice
func NewTodoService(repository repository.TodoRepository) TodoService {
	return TodoService{
		repository: repository,
	}
}
