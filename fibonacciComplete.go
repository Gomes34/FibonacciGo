package main

import "fmt"

var (
	termos = 50
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	for i := 0; i < termos; i++ {
		fmt.Printf("F%d: %d\n", i+1, fibonacci(i))
	}
}
