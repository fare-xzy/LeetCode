package TutorialEasy

import "fmt"

func Variable() {
	f := "Runoob" // var f string = "Runoob"
	fmt.Println(f)

	vname1, vname2, vname3 := "v1", "v2", "v3" // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

	fmt.Println(vname1, vname2, vname3)

}
