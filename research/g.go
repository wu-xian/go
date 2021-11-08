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
}
