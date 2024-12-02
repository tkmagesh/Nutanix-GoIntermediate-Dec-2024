package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		ch := make(chan int)
		data := <-ch
		ch <- 100
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		ch <- 100
		data := <-ch
		fmt.Println(data)
	*/

	/*
		ch := make(chan int)
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		data := <-ch
		fmt.Println(data)
	}()
	ch <- 100
	wg.Wait()
}
