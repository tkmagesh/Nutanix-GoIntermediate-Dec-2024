package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	ch := genNos(stopCh)

	go func() {
		fmt.Println("Hit ENTER to stop....")
		fmt.Scanln()
		// stopCh <-struct{}{}
		close(stopCh)
	}()

	for no := range ch {
		fmt.Println(no)
	}
}

func genNos(stopCh chan struct{}) <-chan int {
	ch := make(chan int)

	go func() {
	LOOP:
		for no := 1; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("Timeout occurred")
				break LOOP
			default:
				ch <- no * 10
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(ch)
	}()
	return ch
}

/*
func genNos() <-chan int {
	ch := make(chan int)
	start := time.Now()
	go func() {
		for no := range 100 {
			time.Sleep(500 * time.Millisecond)
			ch <- (no + 1) * 10
			if elapsed := time.Since(start); elapsed >= 10*time.Second {
				fmt.Println("Time elapsed")
				break
			}
		}
		close(ch)
	}()
	return ch
}
*/
