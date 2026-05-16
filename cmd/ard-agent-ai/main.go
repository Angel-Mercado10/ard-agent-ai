package main

import (
	"fmt"
	"os"

	"github.com/gentleman-programming/ard-agent-ai/internal/app"
)

var version = "dev"

func main() {
	app.Version = version
	if err := app.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
