package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	fmt.Println("Process ID :", os.Getpid())
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	ch := genNos(stopCh)
	for no := range ch {
		fmt.Println(no)
	}
}

func genNos(stopCh chan os.Signal) <-chan int {
	ch := make(chan int)

	go func() {
	LOOP:
		for no := 1; ; no++ {
			select {
			case <-stopCh:
				fmt.Println("Stop signal received.. shutting down")
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
