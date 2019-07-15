package testssd

import (
	"fmt"
	"runtime"
	"os"
)

/*
sx1 := 1
sx2 := - 20
ss := flase
*/

var (
	sx1 uint = 1
	sx2 int = -20
	sx4 int = -4
)
var ss bool = false

func LocalGloab(){
	
	fmt.Printf("gloab S2&  : %v \n", &sx2)
	sx2 := 2
	fmt.Printf("local S2&  : %v \n", &sx2)
}


func FebP2 (k int){
	fmt.Printf("febs start :  ")
	n, m := 0,1
	f := func(a int,b int)(int, int){
			a , b = b , b+a
			return a, b
	}
	for i := 1; i < k; i++ {
		n, m = f(n, m)
		fmt.Printf("%v ", n)
	}
	fmt.Printf("febs2 ends \n")
}

func FebP3 (k int){
	fmt.Printf("febs start :  ")
	n, m := 0,1
	for i := 1; i < k; i++ {
		n, m = func(a int,b int)(int, int){
			a , b = b , b+a
			return a, b
		}(n, m)
		fmt.Printf("%v ", n)
	}
	fmt.Printf(" febs3 ends \n")
	
}

func ViewSys(){
	var goos string = runtime.GOOS
	fmt.Printf("os is %s \n ", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path is %s \n ", path)
}

