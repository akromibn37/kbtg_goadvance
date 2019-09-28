package main

import "fmt"

func main() {
	fruits := []string{"Apple", "Orange", "Plum", "Grape"}
	fmt.Println("o fruits:", fruits)
	var some = make([]string, 3)
	n := copy(some, fruits[1:3:3])
	some = append(some, "Banana")

	fmt.Println("cap some:", n)
	fmt.Println("some:", some)
	fmt.Println("n fruits:", fruits)
	// fmt.Println("some:", some)

}
