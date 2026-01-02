package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/73bits/todo/internal/service"
)

func Edit(svc *service.TodoService) {
	fs := flag.NewFlagSet("edit", flag.ExitOnError)
	index := fs.Int("i", -1, "todo index")
	title := fs.String("title", "", "new title")
	fs.Parse(os.Args[2:])

	if *index < 0 || *title == "" {
		fmt.Fprintln(os.Stderr, "index and title are required")
		os.Exit(1)
	}

	if err := svc.Edit(*index, *title); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("--- todo updated successfully ---")
}
