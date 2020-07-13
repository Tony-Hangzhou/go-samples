package grammar

import (
	"bytes"
	"strings"
)

func String() {

	// 空字符串
	var s string
	println(s == "")
	// 跨行时加号在最后
	s1 := "ab" +
		"cd"
	println(s1 == "abcd")
	// 截取字符串
	println(s1[0:2])
	// 高性能连接字符串
	println(testJoin())
	// 高性能连接字符串
	println(testBuffer())

}

// 分配足够的内存，避免中途扩张底层数组
func testJoin() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}

//
func testBuffer() string {
	var b bytes.Buffer //定义一个缓冲byte类型的缓冲器存放着都是byte
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()

}
