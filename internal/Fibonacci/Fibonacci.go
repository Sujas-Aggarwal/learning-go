package Fibonacci

import (
	"fmt"
)

type Method int

const (
	Recursive Method = iota
	Iterative
)

// basically we will print fibonacci numbers using this upto the numbers the user provides, easy peasy
// the logic for fibonacci is very simple again, just keep 2 numbers, the next one is always the sum of the two
func Fibonacci(n int, algo Method) {
	if n <= 0 {
		return
	}

	switch algo {
	case Recursive:
		for i := 0; i < n; i++ {
			fmt.Println(recursiveFib(i))
		}
	case Iterative:
		iterativeFib(n)
	}
}

func iterativeFib(n int) {
	a := 0
	b := 1
	for range n {
		fmt.Println(a)
		a, b = b, a+b
	}
}

func recursiveFib(n int) int {
	if n <= 1 {
		return n
	}

	return recursiveFib(n-1) + recursiveFib(n-2)
}
