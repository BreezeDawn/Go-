package main

import "fmt"

func main() {
	// 1. channel
	goChannel()

	// 2. channel add
	pipe := make(chan int, 1)
	go goChannelAdd(1, 2, pipe)
	addRes := <-pipe
	fmt.Println(addRes)

	// 3. func return many
	a, b, c := goReturnMany()
	fmt.Println(a, b, c)
}
