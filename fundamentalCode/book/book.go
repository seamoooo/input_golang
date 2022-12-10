package book

type Book struct {
	Title  string
	Author string
	Price  int
}

func NewAuthorlessBook(titile string) *Book {
	return &Book{
		Title:  titile,
		Author: "Unkown Author",
		Price:  0,
	}
}
