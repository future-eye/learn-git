package main

import (
	"fmt"
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch <- res
}

func main() {
	var a int
	for i := 1; i <= 20; i++ {
		go factorial(i)
		a = <-ch
		fmt.Printf("myres[%d] = %d\n",i,a )
	}
}