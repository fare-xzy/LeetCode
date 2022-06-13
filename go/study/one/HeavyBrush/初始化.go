package HeavyBrush

import "fmt"

func init() {
	fmt.Println("默认初始化函数") // 不调用也会自动执行
}

func TestInit() {
	fmt.Println("我是主要的")
}
