package main

import (
	"fmt"
)

func main() {
	r1 := Rock{}
	p1 := Profe{Content: "123321"}
	pp([]Rock{r1},p1)
}

type Rocker interface{
	Puuuuu() Rocker
}

type Rock struct{}
func (h Rock) Puuuuu() Rocker {
	fmt.Println("pooo~")
	return h
}

type Profer interface{
	Poler() string
}
type Profe struct{
	Content string
}
func (h Profe) Poler() string{
	return h.Content
}

func pp[T Rocker,T2 Profer](aa []T,bb T2){
	for _, a := range aa {
		fmt.Println("asddas",a.Puuuuu())
	}
	fmt.Println("poler",bb.Poler())
}
