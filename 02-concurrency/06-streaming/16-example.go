package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for {
		if data, isOpen := <-ch; isOpen {
			time.Sleep(500 * time.Millisecond)
			fmt.Println(data)
			continue
		}
		fmt.Println("All the data have been received")
		break
	}
}

func genNos() chan int {
	ch := make(chan int)
	go func() {
		count := rand.Intn(20)
		for i := range count {
			ch <- (i + 1) * 10
		}
		fmt.Println("[genNos] all the data produced")
		close(ch)
	}()
	return ch
}
