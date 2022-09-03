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
}
