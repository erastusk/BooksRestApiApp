package model

type Book struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Author string  `json:"author"`
}
