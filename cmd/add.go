package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/73bits/todo/internal/service"
)

func Add(svc *service.TodoService) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	title := fs.String("title", "", "todo title")
	fs.Parse(os.Args[2:])

	if *title == "" {
		fmt.Fprintln(os.Stderr, "title is required")
		os.Exit(1)
	}

	if err := svc.Add(*title); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("--- todo added successfully ---")
}
