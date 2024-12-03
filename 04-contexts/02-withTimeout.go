package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	c1Ctx, cancel := context.WithTimeout(rootCtx, 10*time.Second)
	go func() {
		fmt.Scanln()
		cancel()
	}()
	wg.Add(1)
	go printNos(c1Ctx, wg)
	wg.Wait()
	fmt.Println("Done")
}

func printNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for no := 0; ; no++ {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.Canceled {
				fmt.Println("[printNos] cancellation signal received - (programmatic)")
			}
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("[printNos] cancellation signal received - (timeout)")
			}
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println(no)
		}

	}
}
