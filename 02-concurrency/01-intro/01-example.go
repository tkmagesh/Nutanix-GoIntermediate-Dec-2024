package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduling the function to be executed through the scheduler
	f2()
	// block the main() fn so that the scheduler can look for other goroutines that are scheduled and schedule them for execution
	// DO NOT use the following approaches

	time.Sleep(1 * time.Second)
	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
