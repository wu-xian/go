package main

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/wu-xian/go/research/mypkg"
)

type AAA struct {
	name  string
	value string
}

var a = AAA{
	name:  "111",
	value: "222",
}

func main() {
	var e interface{} = a
	ptr := unsafe.Pointer(&e)
	ef := (*runtime.TheEface)(unsafe.Pointer(ptr))
	fmt.Printf("%+v", ef)
	typeIns := ef.GetType()
	fmt.Printf("\n\n%+v\n\n", typeIns)
	fmt.Printf("\ntypename:%+v\n", ef.GetTypeName())

	dataPtr := ef.GetData()
	var aaa *AAA
	aaa = (*AAA)(unsafe.Pointer(dataPtr))
	fmt.Printf("\naaa:%+v", aaa)

	fmt.Println("=============")

	var pp interface{} = mypkg.PrintThis{
		K: "KKK",
		V: "VVV",
		Next: &mypkg.PrintThis{
			K: "KK",
			V: "VV",
			Next: &mypkg.PrintThis{
				K: "K",
				V: "V",
			},
		},
	}

	ppEface := (*runtime.TheEface)(unsafe.Pointer(&pp))
	fmt.Printf("\n%+v\n", ppEface.GetTypeName())
}
