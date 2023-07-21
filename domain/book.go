package domain

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
}
