package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/73bits/todo/internal/service"
)

func Toggle(svc *service.TodoService) {
	fs := flag.NewFlagSet("toggle", flag.ExitOnError)
	index := fs.Int("i", -1, "todo index")
	fs.Parse(os.Args[2:])

	if *index < 0 {
		fmt.Fprintln(os.Stderr, "index is required")
		os.Exit(1)
	}

	if err := svc.Toggle(*index); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println("--- todo toggled successfully ---")
}
