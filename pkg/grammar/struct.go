package grammar

import "fmt"

type user struct {
	name string
	age  int8
}

func (u user) ToString() string { // user is received ，扩展user结构体，相当于javascript Person.prototype.log = function () {)

	return fmt.Sprintf("%s %d", u.name, u.age)

}

type manager struct {
	user  // 匿名实现继承 user
	title string
}

func Struct() {

	var m manager
	m.name = "Tony"
	m.age = 18
	m.title = "Tony Title"

	println(m.ToString())

}
