package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	var divisor int
	for {
		fmt.Println("Enter the divisor")
		fmt.Scanln(&divisor)
		if q, r, err := divideAdapter(100, divisor); err != nil {
			if err == ErrDivideByZero {
				fmt.Println("Please do not attempt to divide by zero. try again!")
				continue
			} else {
				fmt.Println("Unknown error :", err)
				break
			}
		} else {
			fmt.Printf("dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
			break
		}
	}
}

// adapter that converts the panic into an error
func divideAdapter(x, y int) (quotient, remainder int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = err.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party api (cannot change the code)
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
