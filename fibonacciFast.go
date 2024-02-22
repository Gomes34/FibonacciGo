package main

import (
	"fmt"
)

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
	termos := 500
	fmt.Printf("F%d: %d\n", termos, fibonacciIterativo(termos))
}
