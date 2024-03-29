package twtg

import (
	"math"
	"fmt"
)

var Pi float64

/*

变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。

这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。

每个源文件都只能包含一个 init 函数。初始化总是以单线程执行，并且按照包的依赖关系顺序执行。

一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。

init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine

如下面这个例子当中的 backend()：

func init(){
	// setup preparations
	go backend()
}

*/

func init(){
	// init() function computes Pi
	Pi = 4 * math.Atan(1) 
	
}


func TestCase0404(){
	
	fmt.Printf("Pi : %g \n " , Pi * 2)
}



