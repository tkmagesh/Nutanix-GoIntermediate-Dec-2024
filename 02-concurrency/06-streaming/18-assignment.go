package main

import "fmt"

// Modify the below program in such a way that the logic for checking a prime number is executed concurrently

func main() {
	var start, end int = 100, 200
	primes := generatePrimes(start, end)
	for _, primeNo := range primes {
		fmt.Printf("Prime No : %d\n", primeNo)
	}

}

func generatePrimes(start, end int) []int {
	var primes []int
	for no := start; no <= end; no++ {
		if isPrime(no) {
			primes = append(primes, no)
		}
	}
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
