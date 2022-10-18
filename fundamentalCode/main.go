package main

import (
	"fmt"
	"sort"
)

func main() {

	type Person struct {
		FristName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"pat", "Payyerson", 37},
		{"tracy", "Prson", 23},
		{"fred", "Pyerson", 14},
	}

	fmt.Println("初期データ")
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println("sort")
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("sortage")
	fmt.Println(people)
}
