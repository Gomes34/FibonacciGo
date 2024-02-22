package main

import "fmt"

var (
	terms int
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Print("Number of terms: ")
	fmt.Scan(&terms)
	for i := 0; i < terms; i++ {
		fmt.Printf("F%d: %d\n", i+1, fibonacci(i))
	}
}
