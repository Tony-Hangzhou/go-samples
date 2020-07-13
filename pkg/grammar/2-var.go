package grammar

import "strconv"

// 转换成汇编代码 go tool objdump -s "main\.main" test

var x = 100 // 全局变量

// 全局变量，组方式多行定义
var (
	x1, y1 int
	x2, y2 = 100, "abc"
	// b := 1 简短模式赋值不能用于全局变量，只能用于函数内部
)

func Var() {

	/**
	【一】简短模式赋值，只能用于函数内部
	a, s := 1, "abc"
	fmt.Printf("a=%d s=%s", a, s)
	**/

	/**
	【二】等号为分隔，左边为变量，右边为赋值

	var a1, s1 = 100, "abc"
	fmt.Printf("a1=%d s1=%s", a1, s1)
	*/

	/**
	 【三】局部变量、全局变量和简短模式赋值退化。
	 在处理函数返回值时，退化赋值允许我们重复使用err变量，想到方便。

	println(&x, x) // 全局变量
	x := "abc" // 简短模式重新定义局部变量，内存地址不同
	println(&x, x)
	//x := "dd" // 简短模式退化为赋值操作失败，因为退化的前提是至少有一个新变量，否则no new variable on left side of :=
	x, y := "dd", 1 // 简短模式退化为赋值操作成功
	fmt.Printf("%s, %d", x, y)
	*/

	/**
	【四】函数内不同作用域变量定义
	*/
	x := 100
	println(&x, x)

	{
		x, y := 200, 300 // 不同作用域，全部是新变量定义
		println(&x, x, y)
	}

	/**
	【五】空标识符 _ 用来临时规避编译器对未使用变和导入包的错误检查
	*/
	m, _ := strconv.Atoi("12") // 忽略err变量
	println(m)

}
