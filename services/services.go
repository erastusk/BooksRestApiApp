package services

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"

	"github.com/erastusk/books/go_http/model"
)

var (
	id     int
	name   string
	price  float32
	author string
	Dbconn *sql.DB
)

const tablename = "books"

func DbInstance() {
	Dbconn = model.SqlConn()
}

func InsertBook(b model.Book) {
	if CheckIfExist(b.Name) {
		// fmt.Printf("Book: %s exists\n", b.Name)
		return
	}
	ins := `
  INSERT INTO books (name,price,author) 
  VALUES($1 ,$2 ,$3)
  `
	_, err := Dbconn.Exec(ins,
		b.Name,
		b.Price,
		b.Author,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllBooks() []model.Book {
	a := []model.Book{}
	getall := `
  SELECT id, name, price, author FROM books
  `
	rows, err := Dbconn.Query(getall)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &price, &author); err != nil {
			log.Fatal(err)
		}
		b := model.Book{
			Id:     id,
			Name:   name,
			Price:  price,
			Author: author,
		}
		a = append(a, b)

	}
	return a
}

func GetBookById(id int) model.Book {
	getid := `
  SELECT name, price, author FROM books WHERE id=$1
  `
	err := Dbconn.QueryRow(getid, id).Scan(&name, &price, &author)
	if err == sql.ErrNoRows {
		log.Printf("Id %d not found", id)
	}
	return model.Book{
		Id:     id,
		Name:   name,
		Price:  price,
		Author: author,
	}
}

func CheckIfExist(n string) bool {
	ifexist := `
  SELECT name FROM books WHERE name=$1
  `
	row := Dbconn.QueryRow(ifexist, n).Scan(&name)
	if row != sql.ErrNoRows {
		return true
	}
	return false
}

func DeleteBook(id int) error {
	dropbook := `
  DELETE FROM books WHERE id=$1
  `
	r, err := Dbconn.Exec(dropbook, id)
	if err != nil {
		return err
	}
	a, _ := r.RowsAffected()
	if a == 0 {
		return errors.Errorf("No rows affected, book not found")
	}

	return nil
}

func CreateTable() {
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
	_, err := Dbconn.Exec(crTable)
	if err != nil {
		log.Fatal(err)
	}
}
