package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := genNos()
	for data := range ch {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(data)
	}
	fmt.Println("All the data have been received")
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
