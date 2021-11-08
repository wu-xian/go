package main

import "fmt"

func main() {
	go func() {
		fmt.Println("in another g")
	}()
	fmt.Println("mem in go")
}
