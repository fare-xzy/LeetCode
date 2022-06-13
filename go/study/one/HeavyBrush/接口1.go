package HeavyBrush

import "fmt"

type People interface {
	ReturnName(x int) string
}

type Student struct {
	Name string
}

// 定义结构体的一个方法。
// 突然发现这个方法同接口People的所有方法(就一个)，此时可直接认为结构体Student实现了接口People
func (s Student) ReturnName(x int) string {
	fmt.Println(x)
	return s.Name
}

func OutInterface() {
	cbs := Student{Name: "咖啡色的羊驼"}

	var a People
	// 因为Students实现了接口所以直接赋值没问题
	// 如果没实现会报错：cannot use cbs (type Student) as type People in assignment:Student does not implement People (missing ReturnName method)
	a = cbs
	name := a.ReturnName(1)
	fmt.Println(name) // 输出"咖啡色的羊驼"
}
