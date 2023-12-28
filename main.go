package main

import (
	"log"
	"net/http"

	"github.com/erastusk/books/go_http/handlers"
	"github.com/erastusk/books/go_http/logging"
	"github.com/erastusk/books/go_http/services"
)

func main() {
	services.DbInstance()
	// services.CreateTable()
	defer services.Dbconn.Close()
	http.HandleFunc("/books", logging.LoggingMiddleware(handlers.WithJwtAuth(handlers.GetAll)))
	http.HandleFunc("/book", handlers.GetById)
	http.HandleFunc("/book/del", handlers.DeleteBook)
	http.HandleFunc("/book/add", handlers.AddBook)
	http.HandleFunc("/jwt", handlers.GetJwt)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
