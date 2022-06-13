package HeavyBrush

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func ChannelsTest() {
	a := make(chan int)           // 整数类型的无缓冲信道
	b := make(chan int, 0)        // 整数类型的无缓冲信道
	c := make(chan *os.File, 100) // 指向文件指针的带缓冲信道

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println("%T,%T,%T", a, b, c)
}

func WaitChannel() {
	a := make(chan int)
	go func() {
		time.Sleep(5000)
		println("do something")
		a <- 1
	}()
	<-a // 会等待通道中传值
	println("End")
}

var sem = make(chan int, 10)

func Serve(queue []int) {
	for req := range queue {
		println(&req)
		sem <- 1
		go func(req int) { // 此方式可实现更换地址赋值
			println()
			println("This is inner", &req)
			<-sem
		}(req)
	}
}
