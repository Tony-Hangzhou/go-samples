package grammar

import "fmt"

func testParam(x *int) { // 形参，是函数定义中的参数，相当于函数的局部变量
	fmt.Printf("形参 pointer: %p , target: %v\n", &x, x)
}

func testParam2(x *int) { // 延长x生命周期，但是x会被分配到堆上，增加GC回收成本
	go func() { print(x) }()
}

func Func() {
	a := 0x100
	p := &a // 实参，是函数调研时所传递的参数，实参是被调用函数的外部对象
	fmt.Printf("实参 pointer: %p, target: %v\n", &p, p)

	fmt.Println("==============")
	testParam(p) // 在函数调用前，会为形参和返回值分配内存空间，并将实参拷贝到形参内存，不管是指针、引用还是其他类型，都是值拷贝传递。

	testParam2(p)
}
