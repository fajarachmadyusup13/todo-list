package httpsvc

import (
	"github.com/fajarachmadyusup13/todo-list/internal/model"
	"github.com/labstack/echo/v4"
)

type HTTPService struct {
	todoUsecase model.TodoUsecase
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// InitRoutes :nodoc:
func (h *HTTPService) InitRoutes(route *echo.Echo) {
	route.POST("/create-todo", h.CreateToDo)
	route.POST("/update-todo", h.UpdateToDoByID)
	route.POST("/delete-todo", h.DeleteToDoByID)
	route.GET("/todo", h.FindAllToDo)
	route.POST("/todo/:id", h.FindToDoByID)
}

// RegisterToDoUsecase :nodoc:
func (h *HTTPService) RegisterToDoUsecase(t model.TodoUsecase) {
	h.todoUsecase = t
}
