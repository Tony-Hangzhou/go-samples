package grammar

import "fmt"

func Point2() {

	a := 10
	p := &a

	/**
	这里的变量a和p 就是这一块连续地址的首地址的别名，就像ip地址和域名的关系，域名只是ip地址的别名而已.
	a存储是10,p存储的是a的地址.
	*/
	fmt.Printf("变量: %d, 变量内存地址：%p \n", a, p)

	/**
	p2 是二级指针，指向p指针的指针
	*/
	p2 := &p
	fmt.Printf("二级指针: %p", p2)

}
