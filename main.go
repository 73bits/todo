package main

import (
	"fmt"
	"os"

	"github.com/73bits/todo/cmd"
	"github.com/73bits/todo/internal/service"
	"github.com/73bits/todo/internal/store"
)

func run() error {
	help := "expected command: add | list | edit | delete | toggle"

	if len(os.Args) < 2 {
		return fmt.Errorf(help)
	}

	str := &store.JSONStore{"data/todo.json"} // create store
	svc := service.New(str) // create service

	switch os.Args[1] {
	case "add":
		cmd.Add(svc)
	case "list":
		cmd.List(svc)
	case "edit":
		cmd.Edit(svc)
	case "delete":
		cmd.Delete(svc)
	case "toggle":
		cmd.Toggle(svc)
	default:
		return fmt.Errorf("unknown command: %s\n%s", os.Args[1], help)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
