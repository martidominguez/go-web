package main

import (
	"app/internal/application"
	"fmt"
)

func main() {
	// app
	// - config
	app := application.NewServerChi("localhost:8080")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
