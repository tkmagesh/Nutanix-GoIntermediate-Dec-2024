package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
	time.Sleep(500 * time.Millisecond)
	fmt.Println(<-ch)
}

func genNos() chan int {
	ch := make(chan int)
	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		ch <- 40
		ch <- 50
	}()
	return ch
}
