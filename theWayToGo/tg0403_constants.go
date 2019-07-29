package twtg


import (
	"fmt"
	//"fmt"
	//"runtime"
	"os"
)


const clearCode uint8 = 0

const Ln2= 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E= 1/Ln2 // this is a precise reciprocal
const Billion = 1e9 // float constant

const (
	Unknown = 0
	Female = 1
	Male = 2
)

const (
	Monday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)


func TestCase0403(){
	
	fmt.Printf("Wednesday is %d \n",  Wednesday)
	
	fmt.Printf("HOME is %s \n ", os.Getenv("HOME"))
	fmt.Printf("USER is %s \n ", os.Getenv("USER"))
	fmt.Printf("GOROOT is %s \n ", os.Getenv("GOROOT"))
}
// iota 自动加1

/*

指针, slices，maps, channel都是引用类型.

当使用赋值语句 r2 = r1 时，只有引用（地址）被复制.

被引用的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间



*/




