package TutorialEasy

import "fmt"

func Pointer() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

const MAX int = 3

func PointerArray() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func PointerPointer() {
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}

func PointerThree() {
	var a int = 1
	var ptr1 *int = &a
	var ptr2 **int = &ptr1
	var ptr3 **(*int) = &ptr2 // 也可以写作：var ptr3 ***int = &ptr2
	// 依次类推
	fmt.Println("a:", a)
	fmt.Println("ptr1", ptr1)
	fmt.Println("ptr2", ptr2)
	fmt.Println("ptr3", ptr3)
	fmt.Println("*ptr1", *ptr1)
	fmt.Println("**ptr2", **ptr2)
	fmt.Println("**(*ptr3)", **(*ptr3)) // 也可以写作：***ptr3
}
