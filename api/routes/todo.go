package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/infrastructure"
)

//TodoRoutes -> struct
type TodoRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	todoController controllers.TodoController
}

//Setup todo routes
func (i TodoRoutes) Setup() {
	i.logger.Zap.Info("Setting up user routes")
	todos := i.router.Gin.Group("/api/todos")
	{
		// todos.GET("", i.todoController.GetAllTodos)
		todos.POST("", i.todoController.CreateTodo)
		// todos.GET("/:id", i.todoController.GetASingleTodo)
		// todos.PUT("/:id", i.todoController.UpdateTodo)
		// todos.DELETE("/:id", i.todoController.DeleteTodo)
	}
}

func NewTodoRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	todoController controllers.TodoController,
) TodoRoutes {
	return TodoRoutes{
		router:         router,
		logger:         logger,
		todoController: todoController,
	}
}
