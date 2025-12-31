package store

import (
	"encoding/json"
	"os"

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
