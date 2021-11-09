package mypkg

import "fmt"

func RunThis() {
	fmt.Println("run this.")
}

type PrintThis struct {
	K    string
	V    string
	Next *PrintThis
}
