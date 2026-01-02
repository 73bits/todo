package cmd

import (
	"fmt"

	"github.com/73bits/todo/internal/service"
)

func List(svc *service.TodoService) {
	todos, err := svc.List()
	if err != nil {
		fmt.Println(err)
	}

	for i, t := range todos {
		status := " "
		if t.Completed {
			status = "x"
		}
		fmt.Printf("[%s] %d. %s\n", status, i, t.Title)
	}
}
