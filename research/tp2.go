package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	//var x I =A{Name:"asddsa"}
	//Foo("123;;;;;")
	var mys MyStruct[string] = MyStruct[string]{
		Val :"123321aaaaa",
		Name : "asddsa",
	}
	Bar(&mys)

	fmt.Println("=====")

	var s EmbededIface = EmbededStruct{
		Val : "ssss",
	}
	s.Method1(555555)
	fmt.Println(s)
}

type MyInt int
type I interface{
	~int|~int64|~int32
}
type MyI interface{
	I | ~string
}

type MyStruct[T MyI] struct{
	Val T
	Name string
}

type Method1[T any] interface{}
type EmbededIface interface{
	Method1(int)
	Method1[string]
}
type EmbededStruct struct{
	Val Method1[string]
}
func(EmbededStruct) Method1(a int){
	fmt.Println(a)
}

func Bar[T *MyStruct[string]](a T){
	bytes,err:=json.Marshal(a)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}


func Foo[T MyI](a T){
	fmt.Println(a)
	ss := 126
	s := T(ss)
	fmt.Println(s)
}