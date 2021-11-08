package main

import (
	"runtime"
	"time"
)

func main() {
	defer func() {
		runtime.GC()
		println("numg", runtime.NumGoroutine())
	}()
	runtime.GOMAXPROCS(10)
	for i := 0; i < 100000; i++ {
		go timewait()
		println("numg", runtime.NumGoroutine())
		// if i%3 == 0 {
		// 	runtime.GC()
		// }
	}
	println("numg", runtime.NumGoroutine())
	println("numg", runtime.NumGoroutine())
	println("numg", runtime.NumGoroutine())
	println("numg", runtime.NumGoroutine())
}

func timewait() {
	<-time.After(time.Microsecond * 5)
}
