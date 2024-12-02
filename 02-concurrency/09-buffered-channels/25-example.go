package main

import "fmt"

func main() {

	/*
		ch := make(chan int, 1)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int, 3)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))

	ch <- 10
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	ch <- 20
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	ch <- 30
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))

	fmt.Println("channel data:", <-ch)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	fmt.Println("channel data:", <-ch)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
	fmt.Println("channel data:", <-ch)
	fmt.Printf("cap(ch) = %d, len(ch) = %d\n", cap(ch), len(ch))
}
