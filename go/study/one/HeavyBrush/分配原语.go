package HeavyBrush

import (
	"fmt"
	"reflect"
)

type School struct {
	className   string
	teacherName string
}

func TestNew() {
	s := new(School)
	s.className = "1"
	s.teacherName = "2"

	var s1 School

	fmt.Println("%T", s)
	fmt.Println("%T", s1)
}

func TestMake() {
	i := make([]int, 10, 100)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println("%T", i)
	fmt.Println(i)
}
