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
		if q, r, err := divide(100, divisor); err != nil {
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

func divide(x, y int) (quotient, remainder int, err error) {
	fmt.Println("[@divide] calculating quotient")
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	fmt.Println("[@divide] calculating remainder")
	remainder = x % y
	return
}
