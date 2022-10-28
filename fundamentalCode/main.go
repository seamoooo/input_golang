package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

type Manager struct {
	Employee // 組み込みによる合成
	Report   []Employee
}

func main() {
	m := Manager{
		Employee: Employee{
			Name: "上杉謙信",
			ID:   "12344",
		},
		Report: []Employee{},
	}

	fmt.Println(m.Employee)
}
