package logging

import (
	"log"
	"net/http"
	"os"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		duration := time.Since(startTime)
		logger := log.New(os.Stdout, "Logger... | ", log.Ldate)
		logger.Println("| Request Method: ", r.Method)
		logger.Println("| Request URL: ", r.URL)
		logger.Println("| Request Duration: ", duration)

		// next(w, r)
	}
}
