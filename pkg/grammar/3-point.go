package grammar

import "fmt"

/**
Go语言中指针非常常用，一定要掌握
Go语言除了map、slice、chan其他都是值传递，引用传递一定要用指针类型
结构体类型定义方法要注意使用指针类型
接口实现方法时，用指针类型实现的接口函数只能算是指针类型实现的，用结构体类型实现的方法也作为是指针类型实现。


*/

type Person struct { //person结构体，包含年龄，名称，车
	age  int
	name string
	car  Car
}

type Car struct { //person名下的车
	name string //车的名字
}

func (person Person) setName(name string) { //设置名字失败
	person.name = name
}

//当使用指针类型定义方法后，结构体类型的变量调用方法时会自动取得该结构体的指针类型并传入方法。
func (person *Person) setNameForPersonPoint(name string) {
	person.name = name
}

//Person类型指针作为入参，调用时需用 &person 取内存地址
func setNamePoint(p *Person, name string) {
	p.name = name
}

func Point() {

	person := Person{}
	fmt.Println(person) //{0  {}}
	person.age = 12
	person.name = "小明"
	person.car = Car{"宝马"}
	fmt.Println(person)

	/**
	赋值失败

	person.setName("小王")
	fmt.Println(person)
	*/

	/**
	赋值成功
	person.setNameForPersonPoint("小王")
	fmt.Println(person)
	*/

	/**
	赋值成功

	var p *Person = &person
	p.setNameForPersonPoint("小王")
	fmt.Println(*p)
	*/

	/**
	   赋值成功

		setNamePoint(&person, "小王")
	   fmt.Println(person)
	*/

}
