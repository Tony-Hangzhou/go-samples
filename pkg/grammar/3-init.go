package grammar

import "fmt"

type data struct {
	x int
	s string
}

func Init() {

	// 初始化数据结构 "," 结尾
	b := data{
		1,
		"abc",
	}
	fmt.Println(b)

	c := data{1, "bcd"} // 初始化数据结构 "}" 结尾
	fmt.Print(c)
}
