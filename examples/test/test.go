package test

import (
	"fmt"
	"unsafe"
)

func main() {
	calcMap()
}

func calcMap() {
	var result = make(map[string]struct{})
	result["a"] = struct{}{}
	result["b"] = struct{}{}
	result["c"] = struct{}{}
	result["d"] = struct{}{}
	fmt.Println(result)
	fmt.Println(unsafe.Sizeof(result))
	a:=result["a"]
	b:=result["b"]
	c:=result["c"]
	d:=result["d"]
	fmt.Printf("%p - %pp - %pp - %p",&a,&b,&c,&d)
}

func calcSize(){
	var result = make([]string,1000000)
	result[3] = "a"
	result[13] = "a"
	fmt.Println(result)
	fmt.Println(unsafe.Sizeof(result))
}