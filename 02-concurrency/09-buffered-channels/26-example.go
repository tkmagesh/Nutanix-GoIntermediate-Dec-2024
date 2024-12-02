package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("error occurred :", e)
		}
	}()

	// Sync
	/*
		result := divide(100, 0)
		fmt.Println("result :", result)
	*/

	// Async

	// Handle error
	/*
		ch, errCh := divideAsync(100, 0)
		select {
		case result := <-ch:
			fmt.Println("result :", result)
		case e := <-errCh:
			fmt.Println("Error occurred :", e)
		}
	*/

	ch, _ := divideAsync(100, 0)
	fmt.Println(<-ch)

}

func divideAsync(x, y int) (<-chan int, <-chan error) {

	ch := make(chan int)
	errCh := make(chan error, 1)

	go func() {
		defer func() {
			if e := recover(); e != nil {
				errCh <- e.(error)
			}
		}()
		time.Sleep(2 * time.Second)
		result := x / y
		ch <- result
	}()

	return ch, errCh
}

func divide(x, y int) int {
	return x / y
}
