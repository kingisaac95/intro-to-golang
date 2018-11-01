package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()

		time.Sleep(3 * time.Second)
		fmt.Println("Hello World!!!")
	}()

	go func() {
		defer waitGrp.Done()

		fmt.Println("Learning Go Concurrency Basics...")
	}()

	waitGrp.Wait()
}
