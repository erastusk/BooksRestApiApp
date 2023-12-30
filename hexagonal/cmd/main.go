package main

import (
	"hexagonal/internal/adapters/handlers"
	"hexagonal/internal/core/app"
	"hexagonal/internal/core/postgres"
	"net/http"
	"log"
)

func main() {
	//DB
	store := postgres.NewRepository()
	svc := app.NewService(store.GetDB())
	handlers := handlers.NewUserHandlers(svc)
	//http.HandleFunc("/books", logging.LoggingMiddleware(handlers.WithJwtAuth(handlers.GetAll)))
	http.HandleFunc("/book", handlers.GetById)
	http.HandleFunc("/book/del", handlers.DeleteBook)
	http.HandleFunc("/book/add", handlers.AddBook)
	//http.HandleFunc("/jwt", handlers.GetJwt)
	log.Fatal(http.ListenAndServe(":3000", nil))

}