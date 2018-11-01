package main

import (
	"fmt"
)

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	// create new variable with pointer
	// var GoFundamentals courseMeta
	// GoFundamentals := new(courseMeta)

	// create new variable with reference
	GoFundamentals := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 4.8,
	}

	fmt.Println("\nGo Fundamentetal cost, is:", GoFundamentals.Author)
	fmt.Println("\nThe course level is:", GoFundamentals.Level)
	fmt.Println("\nRate:", GoFundamentals.Rating)
}
