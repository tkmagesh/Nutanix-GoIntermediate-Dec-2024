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
	c1Ctx, cancel := context.WithCancel(rootCtx)
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
			fmt.Println("[printNos] cancellation signal received")
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println(no)
		}

	}
}
