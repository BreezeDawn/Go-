package main

import (
	"fmt"
	"time"
)

func goChannel() {
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

func goChannelAdd(a int, b int, pipe chan int) {
	pipe <- a + b
	time.Sleep(5 * time.Second)
}

func goReturnMany() (int, int, int) {
	return 1, 2, 3
}
