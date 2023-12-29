package handlers

import (
	"encoding/json"
	"fmt"
	"hexagonal/internal/core/app"
	"hexagonal/internal/core/domain"
	"io"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	app *app.Service
}

func NewUserHandlers(apps *app.Service) *UserHandlers {
	return &UserHandlers{
		app: apps,
	}
}

func (h *UserHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		res := h.app.GetAllBooks()
		json.NewEncoder(w).Encode(res)
	} else {
		fmt.Fprintf(w, http.ErrNotSupported.ErrorString)
	}
}
func (h *UserHandlers) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		re := domain.Request{}
		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &re)
		h.app.InsertBook(re)
	} else {
		fmt.Fprintf(w, http.ErrNotSupported.ErrorString)
	}
}

func (h *UserHandlers) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	idn, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	res := h.app.GetBookById(idn)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandlers) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		id := r.URL.Query().Get("id")
		idn, err := strconv.Atoi(id)
		if err != nil {
			fmt.Printf("%+v", err)
		}
		err = h.app.DeleteBook(idn)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			fmt.Fprintf(w, "Unable to Delete Book with id: %s\n, not Found", id)
		} else {
			fmt.Fprintf(w, "Deleted Book with id: %s successfully\n", id)
		}
	} else {
		fmt.Fprintf(w, http.ErrNotSupported.ErrorString)
	}
}

// func WithJwtAuth(hfunc http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		logger := log.New(os.Stdout, "Logger... | ", log.Ldate)
// 		logger.Println("| Calling JWT Validation")
// 		tokenString := r.Header.Get("x-jwt-token")
// 		_, err := services.ValidateToken(tokenString)
// 		if err != nil {
// 			fmt.Fprintf(w, "Invalid Token\n")
// 			return
// 		}
// 		hfunc(w, r)
// 	}
// }

// func GetJwt(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		token := services.CreateToken()
// 		fmt.Fprintf(w, token)
// 	} else {
// 		fmt.Fprintf(w, http.ErrNotSupported.ErrorString)
// 	}
// }
