package main

import (
	"os"

	app "github.com/lmfuture-ma/lmaker/example/todolist/cmd"
)

func main() {
	cmd := app.NewTodolistCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
