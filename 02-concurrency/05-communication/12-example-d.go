package main

import (
	"fmt"
	"time"
)

// consumer
func main() {
	ch := add(100, 200)
	/*
		go func() {
			ch <- 10000 // NOT possible in a receive only channel
		}()
	*/
	result := <-ch
	fmt.Println(result)
}

// producer returning a "RECEIVE ONLY" channel
func add(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		result := x + y
		ch <- result
	}()
	return ch
}
