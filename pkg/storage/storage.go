package storage

import (
	"blaze/pkg/models"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type TodoStorage struct {
	client *sqlx.DB
}

func NewTodoStorage(dbUrl string) (*TodoStorage, error) {
	client, err := sqlx.Open("sqlite3", dbUrl)
	if err != nil {
		return nil, err
	}

	return &TodoStorage{client: client}, nil
}

func (s *TodoStorage) GetTodos() ([]models.Todo, error) {
	query := `SELECT id, title, completed FROM todos`

	todos := []models.Todo{}
	err := s.client.Select(&todos, query)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoStorage) GetTodoById(id string) (models.Todo, error) {
	query := `
		SELECT id, title, completed 
		FROM todos WHERE id = ?
	`

	todo := models.Todo{}
	err := s.client.Get(&todo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, ErrTodoNotFound
		}
		return models.Todo{}, err
	}
	return todo, nil
}

func (s *TodoStorage) CreateTodo(newTodo models.Todo) error {
	query := `
		INSERT INTO todos (id, title, completed) 
		VALUES (:id, :title, :completed)
	`

	_, err := s.client.NamedExec(query, &newTodo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoStorage) UpdateTodoStatus(id string, completed bool) error {
	query := `UPDATE todos SET completed = ? WHERE id = ?`

	_, err := s.client.Exec(query, completed, id)
	return err
}

func (s *TodoStorage) DeleteTodoById(id string) error {
	query := `DELETE FROM todos WHERE id = ?`

	_, err := s.client.Exec(query, id)
	return err
}
