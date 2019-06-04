package main

import "fmt"

func go_channel() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	p1 := <-pipe
	fmt.Println(len(pipe))
	pipe <- 4
	fmt.Println(p1)
	fmt.Println(len(pipe))
}
