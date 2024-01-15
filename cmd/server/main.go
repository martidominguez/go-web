package main

import (
	"app/internal/application"
	"fmt"
	"os"
)

func main() {
	// app
	// - config
	token := os.Getenv("API_TOKEN")
	app := application.NewDefaultHTTP(":8080", token)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
