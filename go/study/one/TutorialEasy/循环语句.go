package TutorialEasy

import "fmt"

func Cycle() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func CycleGoto() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
}
