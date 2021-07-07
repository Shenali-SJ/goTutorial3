package mapTest

import "fmt"

var StudentIds = make(map[int]string)
var NilMap map[int]string
var Students = map[int]string{
	1 : "Shen",
	2 : "Thev",
	3 : "Sandu",
}

// AddItems declaration and then adding items
func AddItems(map1 map[int]string) {
	fmt.Println(map1)

	//adding items
	map1[234] = "John"
	map1[123] = "Peter"
	map1[420] = "Anne"
	map1[89] = "Kyle"

	fmt.Println(map1)
}

func AddItems2() {
	employees := map[string]int{
		"Mike" : 1234,
		"Misty" : 563,
	}
	employees["Morgan"] = 3000
	fmt.Println(employees)
}

func FindItem(index int) {
	value, ok := Students[index]
	if ok == true {
		fmt.Println("Value of ", index, " is ", value)
		return
	}
	fmt.Println(index ," is not found")
}


