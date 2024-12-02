package main

import (
	"fmt"
	"sync"
)

// Modify the below program in such a way that the logic for checking a prime number is executed concurrently

func main() {
	var start, end int = 100, 2000
	primes := generatePrimes(start, end)
	for primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
}

func generatePrimes(start, end int) <-chan int {
	primes := make(chan int)
	wg := sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isPrime(no) {
				primes <- no
			}
		}()
	}
	go func() {
		wg.Wait()
		close(primes)
	}()
	return primes
}

func isPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
