package TutorialEasy

import "fmt"

func DataType() {
	var a bool = true // 布尔型
	var b int = 1     //整型
	var b1 int8 = 1   // 有符号 8 位整型 (-128 到 127)
	var b2 int16 = 1  // 有符号16位整型(-32768 到 32767)
	var b3 int32 = 1  // 有符号 32 位整型 (-2147483648 到 2147483647)
	var b4 int64 = 1  //有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	var b5 uint8 = 1  //无符号 8 位整型 (0 到 255)
	var b6 uint16 = 1 //无符号 16 位整型 (0 到 65535)
	var b7 uint32 = 1 //无符号 32 位整型 (0 到 4294967295)
	var b8 uint64 = 1 //无符号 64 位整型 (0 到 18446744073709551615)
	var c float32     // 浮点类型 32位浮点数
	var d float64     // 浮点类型 64位浮点数 精度比32位高，尽量使用64位浮点数
	//z := x + yi
	//x = real(z) 实部
	//y = imag(z) 虚部
	var e complex64 = 1 //用于复数计算，内置real 和 imag 分别用于返回复数的实部和虚部

	var f = "1" // 字符串不可改变

	fmt.Print(a, b, c, d, e, f)
	fmt.Print(b1, b2, b3, b4, b5, b6, b7, b8)
}
