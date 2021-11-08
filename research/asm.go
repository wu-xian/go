package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var c chan int = make(chan int)
	defer func() {
		println("lenc", len(c))
	}()
	go func() {
		c <- add(1, 2)
	}()
	go func() {
		c <- add(2, 3)
	}()
	go func() {
		for c1 := range c {
			println(c1)
		}
	}()
	<-time.After(time.Second)
}

func add(a, b int) int {
	return a + b
}
