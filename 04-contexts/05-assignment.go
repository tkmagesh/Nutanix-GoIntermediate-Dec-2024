/*
use context api for the below functionality
incorporate a timeout based cancellation as well (15 seconds)
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	dataCh := genNos(cancelCtx)
	go func() {
		fmt.Println("Hit ENTER to stop...!")
		fmt.Scanln()
		cancel()
	}()
	for no := range dataCh {
		fmt.Println(no)
	}
	fmt.Println("Done!")
}

func genNos(ctx context.Context) <-chan int {
	dataCh := make(chan int)
	go func() {
	LOOP:
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				break LOOP
			case dataCh <- i:
				time.Sleep(500 * time.Millisecond)
			}
		}
		close(dataCh)
	}()
	return dataCh
}
