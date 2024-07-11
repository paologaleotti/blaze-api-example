package storage

import (
	"blaze/pkg/models"
	"context"
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

func (s *TodoStorage) GetTodos(ctx context.Context) ([]models.Todo, error) {
	query := `SELECT id, title, completed FROM todos`

	todos := []models.Todo{}
	err := s.client.SelectContext(ctx, &todos, query)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *TodoStorage) GetTodoById(ctx context.Context, id string) (models.Todo, error) {
	query := `
		SELECT id, title, completed 
		FROM todos WHERE id = ?
	`

	todo := models.Todo{}
	err := s.client.GetContext(ctx, &todo, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Todo{}, ErrTodoNotFound
		}
		return models.Todo{}, err
	}
	return todo, nil
}

func (s *TodoStorage) CreateTodo(ctx context.Context, newTodo models.Todo) error {
	query := `
		INSERT INTO todos (id, title, completed) 
		VALUES (:id, :title, :completed)
	`

	_, err := s.client.NamedExecContext(ctx, query, &newTodo)
	if err != nil {
		return err
	}
	return nil
}

func (s *TodoStorage) UpdateTodoStatus(ctx context.Context, id string, completed bool) error {
	query := `UPDATE todos SET completed = ? WHERE id = ?`

	_, err := s.client.ExecContext(ctx, query, completed, id)
	return err
}

func (s *TodoStorage) DeleteTodoById(ctx context.Context, id string) error {
	query := `DELETE FROM todos WHERE id = ?`

	_, err := s.client.ExecContext(ctx, query, id)
	return err
}
