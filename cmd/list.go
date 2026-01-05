package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

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

func ListTabWriter(svc *service.TodoService) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	defer w.Flush()

	todos, err := svc.List()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintln(w, "ID\tTITLE\tCOMPLETED\tCREATED_AT\tCOMPLETED_AT")
	for i, t := range todos {
		fmt.Fprintln(w, i, "\t", t.Title, "\t", t.Completed, "\t", t.CreatedAt, "\t", t.CompletedAt)
	}
}
