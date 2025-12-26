package main

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("invalid index")
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tTITLE\tCOMPLETED\tCREATED_AT\tCOMPLETED_AT")
	for index, t := range *todos {
		completedAt := "-"
		if t.CompletedAt != nil {
			completedAt = t.CompletedAt.Format(time.RFC1123)
		}
		fmt.Fprintln(w, index, "\t", t.Title, "\t", t.Completed, "\t", t.CreatedAt.Format(time.RFC1123), "\t", completedAt)
	}
}
