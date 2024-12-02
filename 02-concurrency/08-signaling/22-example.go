package main

import (
	"fmt"
	"time"
)

func main() {
	ch := genNos()
	for no := range ch {
		fmt.Println(no)
	}
}

func genNos() <-chan int {
	ch := make(chan int)
	timeoutCh := timeout(10 * time.Second)
	go func() {
	LOOP:
		for no := 1; ; no++ {
			select {
			case <-timeoutCh:
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

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
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
