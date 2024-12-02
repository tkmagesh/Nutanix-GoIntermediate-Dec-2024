package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		fmt.Println("[@main - deferred]")
		if err := recover(); err != nil {
			fmt.Println("shutting down coz of a panic, err :", err)
			return
		}
		fmt.Println("Done")
	}()
	divisor := 7
	q, r := divide(100, divisor)
	fmt.Printf("dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divide(x, y int) (quotient, remainder int) {
	fmt.Println("[@divide] calculating quotient")
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	fmt.Println("[@divide] calculating remainder")
	remainder = x % y
	return
}
