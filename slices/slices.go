package main

import (
	"fmt"
)

func main() {
	myCourses := make([]string, 5, 10) // <type>, <len>, <cap>
	fmt.Printf("Length: %d \nCapacity: %d", len(myCourses), cap(myCourses))

	// shorthand
	todo := []string{"Eat", "Sleep", "Code"}
	fmt.Printf("\n\nLength: %d \nCapacity: %d\n", len(todo), cap(todo))
}
