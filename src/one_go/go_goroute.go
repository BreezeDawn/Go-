package main

import (
	"fmt"
	"time"
)

func test_goroute(i int) {
	fmt.Println(i)
}

func go_groute() {
	for i := 0; i < 100; i++ {
		go test_goroute(i)
	}
	time.Sleep(time.Second)
}
