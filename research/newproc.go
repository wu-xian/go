package main

import (
	"fmt"
	"time"
)

func main() {
	newproc()
	<-time.After(1 * time.Second)
}

func newproc() {
	go func(a, b int) {
		a = a + b
		fmt.Println(a)
	}(2, 3)
}
