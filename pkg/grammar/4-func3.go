package grammar

import "fmt"

func Func3() {

	/**
	匿名函数直接调用
	*/
	func(s string) {
		fmt.Print(s)
	}("hello")

	/**
	匿名函数赋值给变量
	*/
	add := func(x int, y int) {
		fmt.Printf("\n %d", x+y)
	}

	add(1, 2)

}
