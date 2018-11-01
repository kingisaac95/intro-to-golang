package main

import (
	"fmt"
	"os"
	"reflect"
)

var (
	name   = os.Getenv("USER") // inferred type reflect.TypeOf(name) will be string
	module = 4.3
)

const dob = "01/01/1990"

func main() {
	// short assignment
	foo := "bar"
	bar := &foo // pointer using & | reference with *
	course := "Go Fundamentals"

	fmt.Println("Name is", name)
	fmt.Println("Module is of type", reflect.TypeOf(module))
	fmt.Println("The address of foo is", bar, "and the value of foo is", *bar)
	fmt.Println("Current course is", course)

	changeCourse(&course) // pass address in memory

	fmt.Println("Course is now changed to", course)

	// printAllEnvs()
}

func changeCourse(course *string) string {
	*course = "Web applications with Go"

	return *course
}

func printAllEnvs() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
