package main

import ("fmt"
)

func main(){
	type a struct{
		A string
		B int
	}
	var c int
	isZero[int](c)
	isZero[a](a{})
	//m := map[string]string{}
	//isZero[map[string]string](m)
}

func isZero[T comparable](item T){
	var zero T
	if  item==zero {
		fmt.Printf("%+v is zero\n",item)
	}else{
		fmt.Printf("%+v is not zero\n",item)
	}
}
