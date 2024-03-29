package twtg

import (
	"math"
	"fmt"
)

func TestCase0402() {
	var ss int = 1269
	var msg error
	var sb uint8
	
	sb, msg = Uint8FromInt(ss)
	if nil == msg {
		fmt.Printf("success :%d\n", sb)
	}else{
		fmt.Printf("error : %s\n", msg)
	}
	
	ss = 127
	sb, msg = Uint8FromInt(ss)
	if nil == msg {
		fmt.Printf("success :%d\n", sb)
	}else{
		fmt.Printf("error : %s\n", msg)
	}
		
}

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxInt8 {
		return uint8(n), nil
	}
   
	return 0, fmt.Errorf("%d out of range", n)
}



/*

同一个程序 只有一个main

同一个文件夹下 只有一个package



*/