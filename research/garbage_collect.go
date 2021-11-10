package main

import (
	"fmt"
	"runtime"
)

func main() {
	var aaa string = "asljfdhakfhakushfkuaehfkushkufhskshfkuhsdkufghks"
	fmt.Println(&aaa)
	_ = aaa
	runtime.GC()
}
