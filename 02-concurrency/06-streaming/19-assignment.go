// implementing workers

package main

import (
	"fmt"
	"sync"
)

// Modify the below program in such a way that the logic for checking a prime number is executed concurrently

func main() {
	var start, end int = 100, 200
	primes := generatePrimes(start, end, 10)
	for primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}
	fmt.Println("Done")
}

func generatePrimes(start, end int, workerCount int) <-chan int {
	primes := make(chan int)
	noCh := make(chan int)

	go func() {
		for no := start; no <= end; no++ {
			noCh <- no
		}
		close(noCh)
	}()

	wg := sync.WaitGroup{}
	for idx := range workerCount {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("Starting worker # %d\n", id)
			defer wg.Done()
			for no := range noCh {
				if isPrime(no) {
					fmt.Printf("Prime # %d found at worker %d\n", no, id)
					primes <- no
				}
			}
			fmt.Printf("Shutting down worker # %d\n", id)
		}(idx + 1)
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
