package main

import (
	"fmt"
)

var terms int

func fibonacciIterativo(n int) int {
	if n <= 1 {
		return n
	}

	fib := make([]int, n+1)
	fib[0], fib[1] = 0, 1

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	fmt.Print("Number of terms: ")
	fmt.Scan(&terms)
	fmt.Printf("F%d: %d\n", terms, fibonacciIterativo(terms))
}
