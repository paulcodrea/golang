package main

import "fmt"

// struct Person with 3 fields
// 2 of them are exported (Name, Age)
// 1 of them is not exported (isAdmin)

//struct is a collection of fields (properties) and methods (functions)
// that are grouped together to form a single unit.
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	isAdmin bool
}

func main() {
	p := Person{
		Name:    "Mike",
		Age:     16,
		isAdmin: false,
	}

	r := Person{
		Name:    "Raj",
		Age:     16,
		isAdmin: false,
	}

	fmt.Println(r)
	fmt.Println(p, p.Name, p.Age, p.isAdmin)
}
