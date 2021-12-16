package main

import "fmt"

func Othermain() {
	c := make(chan int, 3)

	c <- 1
	c <- 2

	close(c)
	fmt.Println(<-c)
}
