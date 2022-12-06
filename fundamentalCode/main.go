package main

import "fmt"

type NationalRoute int

const (
	Nkaigo NationalRoute = 300
	AKaido NationalRoute = 301
	Bkaido NationalRoute = 302
)

func (n NationalRoute) String() string {
	switch n {
	case Nkaigo:
		return "長崎"
	case AKaido:
		return "阿蘇"
	case Bkaido:
		return "盆地"
	default:
		return fmt.Sprintf("国体道路")
	}
}

func main() {
	var test NationalRoute = 300

	fmt.Println(test)
}
