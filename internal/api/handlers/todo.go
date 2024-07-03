package handlers

import (
	"blaze/pkg/httpcore"
	"blaze/pkg/models"
	"blaze/pkg/storage"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (c *ApiController) GetTodos(w http.ResponseWriter, r *http.Request) (any, int) {
	todos, err := c.storage.GetTodos()
	if err != nil {
		return httpcore.ErrUnkownInternal.With(err), http.StatusInternalServerError
	}
	return todos, http.StatusOK
}

func (c *ApiController) GetTodo(w http.ResponseWriter, r *http.Request) (any, int) {
	id := r.PathValue("id")

	todo, err := c.storage.GetTodoById(id)
	if err != nil {
		if errors.Is(err, storage.ErrTodoNotFound) {
			return httpcore.ErrNotFound, http.StatusNotFound
		}
		return httpcore.ErrUnkownInternal.With(err), http.StatusInternalServerError
	}

	return todo, http.StatusOK
}

func (c *ApiController) CreateTodo(w http.ResponseWriter, r *http.Request) (any, int) {
	newTodo, err := httpcore.DecodeBody[models.NewTodo](w, r)
	if err != nil {
		return httpcore.ErrBadRequest.With(err), http.StatusBadRequest
	}

	todo := models.Todo{
		Id:        uuid.New().String(),
		Title:     newTodo.Title,
		Completed: false,
	}

	err = c.storage.CreateTodo(todo)
	if err != nil {
		return httpcore.ErrUnkownInternal.With(err), http.StatusInternalServerError
	}

	return todo, http.StatusCreated
}
