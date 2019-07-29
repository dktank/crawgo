package twtg

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

/*

初识Go, 了解 变量 函数 闭包 等简单概念

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

func TestCase0401() {

	LocalGloab()
	ViewSys()

	var (
		da = 9
		//db = 8
		dc = 19
	)
	
	FebP2(da)
	FebP3(dc)

}



/*

25个 关键保留字

break
case
chan
const
continue

default
defer
else
fallthrough
for

func
go
goto
if
import

interface
map
package
range
return

select
struct
switch
type
var







预定义标识符 ( 内置函数 和 基本类型名称 )





append
copy
print
println
recove
panic

make 
close
string
iota 
new
nil

imag
real
cap
uintptr
byte
len

true
false
bool


int  int8 int16 int32 int64

uint uint8 uint16 uint32 uint64

float32 float64

complex complex64 complex128

分隔符

()  []  {}

标点

, .  ; : ...




*/
