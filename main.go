package main

import (
	"fmt"
)

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) (returns int) {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Hello from main.go")
	n := 3

	fmt.Printf("Factorial of %d  is %d\n", n, factorial(n))

	i := 10
	// for i := 0; i < 10; i++ {
	fmt.Printf("Fibonacci of %d  is %d\n", i, fibonacci(i))
	// }
}
