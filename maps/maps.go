package main

import (
	"fmt"
)

func main() {
	// createMaps()
	printData()
}

func createMaps() {
	myCourses := make(map[string]int)
	myCourses["CS-101"] = 4
	myCourses["MTH-111"] = 3
	myCourses["STA-111"] = 3

	// shorthand
	studentBio := map[string]string{
		"name":    "Kingdom Orjiewuru",
		"favFood": "Fried Rice",
	}
	fmt.Printf("\nMy Courses: %v\nBio: %v\n", myCourses, studentBio)
}

// this is basically to test how go randomizes how values are returned in maps
// called random offset
func printData() {
	testMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
	}
	testMap["F"] = 120

	for key, value := range testMap {
		fmt.Printf("Key: %v Value: %v\n", key, value)
	}

	delete(testMap, "F") // delete map item

	fmt.Println(testMap)
}
