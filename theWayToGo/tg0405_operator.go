package twtg

/*

一元运算符用于一个值的操作（作为后缀）
二元运算符则可以和两个值或者操作数结合（作为中缀）

两个类型相同的值才可以和二元运算符结合
Go是强类型语言，不进行隐式转换，
任何不同类型之间的转换都必须显式说明

Go 不存在像 C 和 Java 那样的运算符重载，表达式的解析顺序是从左至右



非 !  
和&& 
或||


格式化输出时，用 %t 表示输出的值为布尔型



对于布尔值的好的命名能够很好地提升代码的可读性，
例如以 is 或者 Is 开头的 isSorted、isFinished、isVisible，
使用这样的命名能够在阅读代码的获得阅读正常语句一样的良好体验，
例如标准库中的 unicode.IsDigit(ch)


基于架构的类型，例如：int、uint 和 uintptr。

float32 精确到小数点后 7 位，
float64 精确到小数点后 15 位。
由于精确度的缘故，在使用 == 或者 != 来比较浮点数时应当非常小心。
你最好在正式使用前测试对于精确度要求较高的运算


%d 用于格式化整数（%x 和 %X 用于格式化 16 进制表示的数字
%g 用于格式化浮点型（%f 输出浮点数，%e 输出科学计数表示法）
%0nd 用于规定输出长度为n的整数，其中开头的数字 0 是必须的。

%n.mg 用于表示数字 n 并精确到小数点后 m 位
除了使用 g 之外，还可以使用 e 或者 f
例如：使用格式化字符串 %5.2e 来输出 3.4 的结果为 3.40e+00。



在使用格式化说明符时，可以使用 %v 来表示复数
但当你希望只表示其中的一个部分的时候需要使用 %f。

逻辑运算符：== != < <= > >=
整数和浮点数的二元运算符有 + - *  /

浮点数除以 0.0 会返回一个无穷尽的结果，使用 +Inf 表示。


语句 b = b + a 简写为 b+=a , -=、*=、/=、%=。

对于整数和浮点数，你可以使用一元运算符 ++（递增）和 --（递减），但只能用于后缀

带有 ++ 和 -- 的只能作为语句，
而非表达式，因此 n = i++ 这种写法是无效的，
其它像 f(i++) 或者 a[i]=b[i++] 
这些可以用于 C、C++ 和 Java 中的写法在 Go 中也是不允许的

在运算时 溢出 不会产生错误，Go 会简单地将超出位数抛弃。
如果你需要范围无限大的整数或者有理数（意味着只被限制于计算机内存），
你可以使用标准库中的 big 包，该包提供了类似 big.Int 和 big.Rat 这样的类型（第 9.4 节）


优先级 	运算符
 7 		^ !
 6 		* / % << >> & &^
 5 		+ - | ^
 4 		== != < <= >= >
 3 		<-
 2 		&&
 1 		||

可以通过使用括号来临时提升某个表达式的整体运算优先级。


别名
当你在使用某个类型时，你可以给它起另一个名字
然后你就可以在你的代码中使用新的名字（用于简化名称或解决名称冲突）。

在 type TZ int 中，TZ 就是 int 类型的新名称（用于表示程序中的时区），
然后就可以使用 TZ 来操作 int 类型的数据



格式化说明符 %c 用于表示字符；
当和字符配合使用时，%v 或 %d 会输出用于表示该字符的整数；
%U 输出格式为 U+hhhh 的字符串（另一个示例见第 5.4.4 节）。


包 unicode 包含了一些针对测试字符的非常有用的函数（其中 ch 代表字符）：
	判断是否为字母：unicode.IsLetter(ch)
	判断是否为数字：unicode.IsDigit(ch)
	判断是否为空白符号：unicode.IsSpace(ch)
这些函数返回一个布尔值。包 utf8 拥有更多与 rune 相关的函数。


*/
import (
	"math"
	"fmt"
	"time"
	"math/rand"
)

const ErrorCode uint8 = 0

func Uint8FromInt_2(n int)(uint8, error){
	if 0 <= n && n <= math.MaxUint8{
		return uint8(n), nil
	}
	return ErrorCode, fmt.Errorf("%d is out of uint8 range", n)
}

func IntFromFloat64(n float64) int{
	// x lies in the integer range
	if math.MinInt32 <= n && n <= math.MaxFloat64{
		whole, fraction := math.Modf(n)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%d is out of int32 range", n))
}

func ComplexTest(){
	var c1 complex64 = 5 + 10i
	fmt.Printf("C1 : %v \n ", c1)
	
}

type ByteSize float64
// 使用位左移与 iota 计数配合可优雅地实现存储单位的常量枚举
const(
	_ = iota  // 站位来忽略
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB 
	
)

//在通讯中使用位左移表示标识的用例
type BitFlag int
const (
	
	// 1 << 0 == 1
	Actice BitFlag = 1 << iota
	// 1 << 1 == 2
	Send
	// 1 << 2 == 4
	Receive
)


func TestRandom(){

	fmt.Printf("Random 5: ")
	for i := 0; i < 5; i++{
		fmt.Printf("%d, ", rand.Int())
	}
	fmt.Printf("\n ")
	
	
	fmt.Printf("Random 5n: ")
	for i := 0; i < 5; i++{
		fmt.Printf("%d, ", rand.Intn(8))
	}
	fmt.Printf("\n ")
	
	/* 
	可以使用 Seed(value) 函数来提供伪随机数的生成种子，
	一般情况下都会使用当前时间的纳秒级数字
	*/
	fmt.Printf("Random 5: ")
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 5; i++{
		fmt.Printf("%2.2f, ", 100 * rand.Float32())
	}
	fmt.Printf("\n ")
}

func TestChar(){
	
	var ch1 byte = 'A'
	
	ch1 = 65
	ch1 = '\x41'
	ch1 = '\377'
	
	
	var ch int = 0
	
	ch = int(ch1)
	
	ch = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)
	
}



func TestCase0405(){
	
	var c1 float64 = 123.623
	var c2 int = 256 
	var c3 int = 119
	
	
	r1, err1 := Uint8FromInt_2(c2)
	r2, err2 := Uint8FromInt_2(c3)
	
	r3 := IntFromFloat64(c1)
	
	fmt.Printf("t1: %d, %s\n", r1, err1)
	fmt.Printf("t2: %d, %s\n", r2, err2)
	fmt.Printf("t3: %d \n ", r3)
	
	ComplexTest()
	/*
	按位补足 ^：
	该运算符与异或运算符一同使用，即 m^x，
	对于无符号 x 使用“全部位设置为 1”，
	对于有符号 x 时使用 m=-1。例如：
	*/
	//  ^10 = -01 ^ 10 = -11
	fmt.Printf("t4: %d  \n", ^10)
	
	/*
	fmt.Printf("math.MinInt8   %d \n", math.MinInt8)
	fmt.Printf("math.MinInt16   %d \n", math.MinInt16)
	fmt.Printf("math.MinInt32   %d \n", math.MinInt32)
	fmt.Printf("math.MinInt64   %d \n", math.MinInt64)
	fmt.Printf("math.MaxInt8   %d \n", math.MaxInt8)
	fmt.Printf("math.MaxInt16   %d \n", math.MaxInt16)
	fmt.Printf("math.MaxUint8   %d \n", math.MaxUint8)
	fmt.Printf("math.MaxFloat64   %g \n", math.MaxFloat64)
	*/
	TestRandom()
	
	TestChar()
	
}

















