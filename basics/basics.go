package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	variables()
	erroHandling()
	maps()
	slices()
	goRoutines()
	references()
}

func variables() {
	var x int
	x = 8
	y := 10
	fmt.Println(x)
	fmt.Println(y)
}

func erroHandling() {
	myVal, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myVal)
	}
}

func maps() {
	m := make(map[string]int)
	m["Hello"] = 2505
	fmt.Println(m["Hello"])
}

func slices() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	for index, value := range s {
		fmt.Println(index, value)
	}
}

func goRoutines() {
	// Go routines
	c := make(chan int)
	go doSomething(c)
	<-c
}

func doSomething(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("doSomething's done")

	c <- 1
}

func references() {
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h, *h)
}
