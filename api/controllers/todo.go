package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	logger      infrastructure.Logger
	todoService services.TodoService
	env         infrastructure.Env
}

func NewTodoController(
	logger infrastructure.Logger,
	todoService services.TodoService,
	env infrastructure.Env,
) TodoController {
	return TodoController{
		logger:      logger,
		todoService: todoService,
		env:         env,
	}
}

func (cc TodoController) CreateTodo(c *gin.Context) {
	todo := models.Todo{}
	// trx := c.MustGet(constant.DBTransaction)

	if err := c.ShouldBind(&todo); err != nil {
		cc.logger.Zap.Error("Error[CreateTodo] (shouldBindJson) : ", err)
		responses.ErrorJSON(c, http.StatusInternalServerError, "Failed To Create User")
		return
	}

	// if err := cc.todoService.WithTrx(trx).CreateTodo(todo); err != nil {

	// }
}
