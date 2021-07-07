package main

import (
	"fmt"
	"goTutorial3/mapTest"
	"goTutorial3/structTest"
)

func main() {

	//variadic
	findNum(12, 8, 3, 54, 90, 10, 12)

	//zero args for variadic parameters
	findNum(3)

	//passing a slice with ellipsis
	nums := []int{12, 43, 78, 90, 11}
	findNum(43, nums...)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)

	fmt.Println()

	//map
	ids := mapTest.StudentIds

	mapTest.AddItems(ids)

	mapTest.AddItems2()

	//nil map
	var map101 = mapTest.NilMap
	fmt.Println("Nil map: ", map101)

	//retrieve values for a given key
	id := 234
	student := ids[id]
	fmt.Println("Student with the id ", id, " is ", student)
	fmt.Println("Student with the id ", 7009, " is ", ids[7009])

	mapTest.FindItem(2)

	//iterate
	for key, value := range ids {
		fmt.Printf("Key: %d, value: %s \n", key, value)
	}

	//delete
	fmt.Println("Before deleting: ", ids)
	delete(ids, 89)
	fmt.Println("After deleting: ", ids)

	//length
	fmt.Println("Length is ", len(ids))

	var word = "Hello"
	//print bytes - utf-8
	for i := 0; i < len(word); i++ {
		fmt.Printf("String %x \n", word[i])
	}

	//print characters
	for i := 0; i < len(word); i++ {
		fmt.Printf("Character %c \n", word[i])
	}

	printChars("Señor")

	fmt.Println()

	//pointers
	var pointer *string = &word
	fmt.Println(pointer)

	size := new(int)
	fmt.Println(size)

	// de referencing a pointer
	*size = 85
	fmt.Println(size, *size)

	fmt.Println()

	//struct
	structTest.CreateEmployee()

	//anonymous struct
	emp := structTest.Emp
	fmt.Println(emp)

	//nested struct
	structTest.NestedStruct()

}

func findNum(findMe int, list...int) {
	found := false
	for i,v := range list {
		if v == findMe {
			fmt.Println("Found ", findMe, " in ", i, "th index")
			found = true
		}
	}
	if !found {
		fmt.Println("Not found")
	}
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println("S: ", s)
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("Chars %c \n", s[i])
	}

	fmt.Println()

	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("Rune %c \n", runes[i])
	}
}
