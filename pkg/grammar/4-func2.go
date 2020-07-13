package grammar

import (
	"fmt"
	"github.com/pkg/errors"
)

type serverOption struct {
	address string
	port    int
	path    string
}

/**
初始化 serverOption 结构体
*/
func newOption() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test"}
}

/**
变参只能放在参数末尾
*/
func testVarParam(s string, a ...int) {
	fmt.Printf("%s, %d", s, a)
}

/**
命名返回值，多参数时可读性强，而且可以return语句隐式返回
*/
func div(x, y int) (z int, err error) {
	if y == 0 {
		err = errors.New("division by zero")
	}
	return
	z = x / y
	return
}

func Func2() {
	o := newOption()
	fmt.Printf("port: %d \n", o.port)
	o.port = 9090
	fmt.Printf("port: %d \n", o.port)

	testVarParam("s", 1, 2, 3)

	div(50, 100)
}
