package main

import (
	"fmt"
)

func main() {
	p1 := Profe{Content: "123321"}
	pp(p1)
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

func pp[T2 Profer](bb T2){
	fmt.Println("poler",bb.Poler())
}
