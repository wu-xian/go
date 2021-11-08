package main

import (
	"runtime"
	"sync"
)

func main() {

	//a, b := gcount1()
	//println("gcount1", a, b)
	c, d := gcount2()
	println("gcount2", c, d)
}

func gcount1() (allgcount, gcount int) {
	defer func() {
		println("-gcount", runtime.GetGCount())
		println("-ggcount", runtime.NumGoroutine())
		allgcount = runtime.GetGCount()
		gcount = runtime.NumGoroutine()
	}()
	runtime.GOMAXPROCS(8)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func(ccc int) {
			var a = 1
			var b = 2
			var d = a
			var e = b
			var _ = d
			var _ = e
			if ccc%3 == 0 {
				runtime.GC()
				println("gcount", runtime.GetGCount())
				println("ggcount", runtime.NumGoroutine())
			}
			println(ccc)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	runtime.GC()
	return
}

func gcount2() (allgcount, gcount int) {
	defer func() {
		println("-gcount", runtime.GetGCount())
		println("-ggcount", runtime.NumGoroutine())
		allgcount = runtime.GetGCount()
		gcount = runtime.NumGoroutine()
	}()
	runtime.GOMAXPROCS(8)
	ch := make(chan struct{}, 10)
	for i := 0; i < 100; i++ {
		go func(ccc int) {
			var a = 1
			var b = 2
			var d = a
			var e = b
			var _ = d
			var _ = e
			if ccc%3 == 0 {
				runtime.GC()
				println("gcount", runtime.GetGCount())
				println("ggcount", runtime.NumGoroutine())
			}
			println(ccc)
			<-ch
		}(i)
		ch <- struct{}{}
	}

	runtime.GC()
	return
}
