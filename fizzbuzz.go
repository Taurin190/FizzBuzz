package main

import (
	"fmt"
)

func printFizzBuzz(num int) {
	if num % 15 == 0 {
		fmt.Println("FizzBuzz")
	} else if num % 5 == 0 {
		fmt.Println("Buzz")
	} else if num % 3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println(num)
	}
}

func main() {
	i := 1
	for i <= 100 {
		printFizzBuzz(i)
		i++
	}
}