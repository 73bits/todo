package store

import "github.com/73bits/todo/internal/model"

type TodoStore interface {
	Load() ([]model.Todo, error)
	Save([]model.Todo) error
}
