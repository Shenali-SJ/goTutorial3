package main

import "fmt"

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
