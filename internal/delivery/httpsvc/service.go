package httpsvc

import (
	"CRUDWithCockroach/internal/model"

	"github.com/labstack/echo/v4"
)

type HTTPService struct {
	todoUsecase model.ToDoUsecase
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
func (h *HTTPService) RegisterToDoUsecase(t model.ToDoUsecase) {
	h.todoUsecase = t
}
