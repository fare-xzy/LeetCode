package TutorialEasy

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func FunctionTest() {
	//func function_name( [parameter list] ) [return_types] {
	//函数体
	//}
	fmt.Println(max(1, 2))
	s, s2 := swap("2", "1")
	fmt.Println(s, s2)
}

func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

//值传递
func swap1(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

//引用传递 传递地址
/* 定义交换值函数*/
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

//函数参数

// 声明一个函数类型
type cb func(int) int

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
func FunctionTest1() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}

//defer 适用于函数结束后的处理，例如文件流的关闭
//defer 执行顺序为倒叙
var mu sync.Mutex

func lock() {
	mu.Lock()
	log.Printf("lock")
}

func unlock() {
	mu.Unlock()
	log.Printf("unlock")
}

func Foo() int {
	lock()

	func() {
		log.Printf("entry inner")
		defer unlock()
		log.Printf("exit inner")
	}()

	time.Sleep(1 * time.Second)
	log.Printf("return")
	return 0
}
