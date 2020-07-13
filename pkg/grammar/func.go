package grammar

func test(x int) func() { //返回函数类型

	return func() { // 匿名函数
		defer println("defer 延迟调用") /*可定义多个defer，按FILO顺序执行*/
		println(x)                  // 闭包
	}

}

func FuncReturn() {

	x := 100
	f := test(x)
	f()

}
