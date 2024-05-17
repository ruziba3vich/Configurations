package models

type Album struct {
	id     int
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a *Album) GetId() int {
	return a.id
}

func (a *Album) SetId(id int) {
	a.id = id
}
