package main

import "fmt"


func fibonacci(n int) []int {
	fib := make([]int, n)

    if n >= 1 {
        fib[0] = 1
    }
    if n >= 2 {
        fib[1] = 1
    }

	for i := 2; i < n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }

	return fib
}

func main(){
	n := 10
	fmt.Println(fibonacci(n))
}