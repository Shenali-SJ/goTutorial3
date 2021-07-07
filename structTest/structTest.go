package structTest

import "fmt"

// Employee exported struct
//exported fields
type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

// Address exported struct
//exported fields
type Address struct {
	City string
	Country string
}

// Person exported struct
//fields are not exported
type Person struct {
	name string
	age int
	Address
}

var Emp = struct {
	name string
	id int
} {
	name : "Nikola",
	id : 45,
}

// CreateEmployee named struct
func CreateEmployee() {
	emp1 := Employee{
		FirstName: "John",
		LastName:  "White",
		Age:       23,
	}
	fmt.Println(emp1)

	//dot operator is used to access individual fields
	fmt.Println("First name: ", emp1.FirstName)

	emp2 := Employee{"Anne", "Marie", 30}
	fmt.Println(emp2)
}

func NestedStruct() {
	p := Person{
		name: "Mia",
		age: 21,
		Address: Address{
			City: "Melbourne",
			Country: "Australia",
		},
	}
	fmt.Println(p)

	//promoted fields
	fmt.Println("City of living: ", p.City)
}