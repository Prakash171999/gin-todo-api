package repository

import (
	"boilerplate-api/infrastructure"
)

// UserRepository -> database structure
type TodoRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewUserRepository -> creates a new User repository
func NewTodoRepository(db infrastructure.Database, logger infrastructure.Logger) TodoRepository {
	return TodoRepository{
		db:     db,
		logger: logger,
	}
}
