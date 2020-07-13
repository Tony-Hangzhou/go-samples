package grammar

import "fmt"

type Printer interface {
	CustomPrint()
}

type Printer2 interface {
	Printer // 迁入其他接口，目标类型方法集中，必须拥有包含潜入接口方法在内的所有方法，才算实现了该接口
	CustomPrint2()
}

type city struct {
	name string
	size int
}

/**
接口 duck type 模式，实现了接口方法的func都实现了接口
*/
func (c city) CustomPrint() {

	fmt.Printf("Tony %s, %d", c.name, c.size)

}

func (c city) CustomPrint2() {

	fmt.Printf("Tony - 2 %s, %d", c.name, c.size)

}

func InterfaceMain() {
	var c city
	c.name = "杭州"
	c.size = 70

	var p Printer2 = c
	p.CustomPrint()
	p.CustomPrint2()

	var o interface{} = c // 空接口，相当于 java.lang.Object 可接收任何类型对象
	if _, ok := o.(Printer2); ok {
		fmt.Println(ok)
	}

	if _, ok := o.(Printer); ok {
		fmt.Println(ok)
	}
}
