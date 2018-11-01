package main

import (
	"fmt"
	"time"
)

func main() {
	// rocketShip()
	findConflicts()
}

func rocketShip() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}
}

func findConflicts() {
	coursesInProg := []string{"Mathematics for Data Science", "Intro to Data Science", "Go Fundamentals"}
	coursesCompleted := []string{"Mathematics for Computer Science", "Go Fundamentals", "Advanced Rocket Science"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("\nHey! We found a conflict.",
					i, "is in both lists!")
			}
		}
	}
}
