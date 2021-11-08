package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := runtime.GetCurG()
	fmt.Printf("%+v", a.G)
	as := runtime.GetAllGs()
	fmt.Printf("%+v", as)
	for i, aa := range as {
		fmt.Println()
		fmt.Println(i)
		fmt.Printf("%+v", aa.G)
		fmt.Println()
	}

	defer func() {
		//defer1
		var a = 1
		var _ = a
	}()

	defer func() {
		//defer2
		var a = 2
		var _ = a
	}()
	fmt.Printf("\n\np:  %+v", runtime.GetCurP().P)

	fmt.Println("gcount:", runtime.GetGCount())
}
