package app

import (
	"database/sql"
	"hexagonal/internal/core/domain"
	"log"

	"github.com/pkg/errors"
	_ "github.com/lib/pq"
)

var (
	id     int
	name   string
	price  float32
	author string
)

type Service struct {
	db *sql.DB
}

func NewService(database *sql.DB) *Service {
	return &Service{
		db: database,
	}
}
func (s *Service) InsertBook(b domain.Request) {
	if s.CheckIfExist(b.Name) {
		// fmt.Printf("Book: %s exists\n", b.Name)
		return
	}
	ins := `
  INSERT INTO books (name,price,author) 
  VALUES($1 ,$2 ,$3)
  `
	_, err := s.db.Exec(ins,
		b.Name,
		b.Price,
		b.Author,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Service) GetAllBooks() []domain.Response {
	a := []domain.Response{}
	getall := `
  SELECT id, name, price, author FROM books
  `
	rows, err := s.db.Query(getall)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &price, &author); err != nil {
			log.Fatal(err)
		}
		b := domain.Response{
			Id:     id,
			Name:   name,
			Price:  price,
			Author: author,
		}
		a = append(a, b)

	}
	return a

}
func (s *Service) GetBookById(id int) domain.Response {
	getid := `
  SELECT name, price, author FROM books WHERE id=$1
  `
	err := s.db.QueryRow(getid, id).Scan(&name, &price, &author)
	if err == sql.ErrNoRows {
		log.Printf("Id %d not found", id)
	}
	return domain.Response{
		Id:     id,
		Name:   name,
		Price:  price,
		Author: author,
	}
}

func (s *Service) CheckIfExist(n string) bool {
	ifexist := `
  SELECT name FROM books WHERE name=$1
  `
	row := s.db.QueryRow(ifexist, n).Scan(&name)
	if row != sql.ErrNoRows {
		return true
	}
	return false
}

func (s *Service) DeleteBook(id int) error {
	dropbook := `
  DELETE FROM books WHERE id=$1
  `
	r, err := s.db.Exec(dropbook, id)
	if err != nil {
		return err
	}
	a, _ := r.RowsAffected()
	if a == 0 {
		return errors.Errorf("No rows affected, book not found")
	}

	return nil
}

func (s *Service) CreateTable() {
	/*
	   - ID
	   - Name
	   - Author
	   - Date Added
	*/

	crTable := `CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6,2) NOT NULL,
    author VARCHAR(100) NOT NULL,
    created timestamp DEFAULT NOW()
  )`
	_, err := s.db.Exec(crTable)
	if err != nil {
		log.Fatal(err)
	}
}
