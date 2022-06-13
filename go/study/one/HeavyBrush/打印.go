package HeavyBrush

import (
	"fmt"
	"os"
	"time"
)

func PrintText() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))

	type T struct {
		a int
		b float64
		c string
	}
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)  // 值输出
	fmt.Printf("%+v\n", t) // 参数名+值
	fmt.Printf("%#v\n", t) // 标准的go版本打印方式
	fmt.Printf("%#v\n", time.RFC822Z)

	fmt.Printf("%T\n", time.RFC822Z) // 打印的是类型
}
