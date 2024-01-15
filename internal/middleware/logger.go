package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logic before
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("Date: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Printf("URL: %s\n", r.URL)
		fmt.Printf("Bytes: %d\n", r.ContentLength)

		// call next handler
		next.ServeHTTP(w, r)
	})
}
