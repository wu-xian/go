package main

import (
	"fmt"
)

func main() {
	Pp([]string{"1","2","3"})
	fmt.Println()
	Pp([]int{1,2,3,4,5,6})
}

func Pp[T any](aa []T){
	for _, a := range aa {
		fmt.Println(a)
	}
}
