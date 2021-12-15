package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	ch1 := make(chan struct{}, 0)
	go func() {
		ch1 <- struct{}{}
	}()
	<-ch1
}
