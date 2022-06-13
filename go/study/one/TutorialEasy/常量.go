package TutorialEasy

import (
	"fmt"
	"unsafe"
)

const (
	a1 = "abc"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

func Constant() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}

func Constant1() {
	println(a1, b1, c1)
}

//iota
//iota，特殊常量，可以认为是一个可以被编译器修改的常量。
//iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
//iota 可以被用作枚举值：
const (
	a2 = iota
	a3
	a4
	a5 = "ha"
	a6
	a7 = iota
	a8
	a9
)

func TestIota() {
	fmt.Println(a2, a3, a4, a5, a6, a7, a8, a9)
}
