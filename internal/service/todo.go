package service

import (
	"errors"
	"time"

	"github.com/73bits/todo/internal/model"
	"github.com/73bits/todo/internal/store"
)

type TodoService struct {
	store store.TodoStore
}

func New(store store.TodoStore) *TodoService {
	return &TodoService{store: store}
}

func (s *TodoService) Add(title string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	todos, err := s.store.Load()
	if err != nil {
		return err
	}

	todos = append(todos, model.Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	})

	return s.store.Save(todos)
}

func (s *TodoService) Delete(index int) error {
	todos, err := s.store.Load()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(todos) {
		return errors.New("invalid todo index")
	}

	todos = append(todos[:index], todos[index+1:]...)
	return s.store.Save(todos)
}

func (s *TodoService) Edit(index int, newTitle string) error {
	if newTitle == "" {
		return errors.New("title cannot be empty")
	}

	todos, err := s.store.Load()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(todos) {
		return errors.New("invalid todo index")
	}

	todos[index].Title = newTitle
	return s.store.Save(todos)
}

func (s *TodoService) Toggle(index int) error {
	todos, err := s.store.Load()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(todos) {
		return errors.New("invalid todo index")
	}

	todo := &todos[index]

	if todo.Completed {
		todo.Completed = false
		todo.CompletedAt = nil
	} else {
		now := time.Now()
		todo.Completed = true
		todo.CompletedAt = &now
	}

	return s.store.Save(todos)
}

func (s *TodoService) List() ([]model.Todo, error) {
	return s.store.Load()
}
