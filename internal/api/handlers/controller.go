package handlers

import (
	"blaze/pkg/storage"
)

type ApiController struct {
	storage *storage.TodoStorage
}

func NewApiController(storage *storage.TodoStorage) *ApiController {
	return &ApiController{
		storage: storage,
	}
}
