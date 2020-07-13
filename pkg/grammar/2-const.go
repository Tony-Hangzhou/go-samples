package grammar

/**
常量是恒定不可变值，通常是字面量
*/
const name = "Lee"

func Const() {

	const sex = "man"
	const age = 12 // const 不适用不会引起编译器错误

	{
		const sex = "woman" // 不同作用域的const

	}

}
