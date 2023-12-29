package domain

type Response struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Author string  `json:"author"`
}

type Request struct {
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Author string  `json:"author"`
}
