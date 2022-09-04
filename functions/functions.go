package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(x)
	fmt.Println(y)

	c := make(chan int)

	go func() {
		fmt.Println("Start")
		time.Sleep(1 * time.Second)
		fmt.Println("Done")

		c <- 1
	}()

	<-c

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	fmt.Println(getValues(1))
}

func sum(values ...int) int {
	total := 0

	for _, num := range values {
		total += num
	}

	return total
}

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x

	return
}
