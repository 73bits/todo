package store

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/73bits/todo/internal/model"
)

type JSONStore struct {
	Path string
}

func (s *JSONStore) Load() ([]model.Todo, error) {
	if _, err := os.Stat(s.Path); os.IsNotExist(err) {
		return []model.Todo{}, nil
	}

	data, err := os.ReadFile(s.Path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []model.Todo{}, nil
	}

	var todos []model.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *JSONStore) Save(todos []model.Todo) error {
	if err := os.MkdirAll(filepath.Dir(s.Path), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	tmp := s.Path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, s.Path)
}
