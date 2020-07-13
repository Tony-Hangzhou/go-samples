package grammar

import "fmt"

func Switch() {

	a, b, c := 1, 2, 3

	switch 1 {
	case a, b: // 相当于或者 ，无须执行break，case语句执行完后自行中断
		fmt.Println("a | b")
		fallthrough //继续执行下一个case，但不再匹配下一个条件,必须是case最后一条语句
	case c:
		fmt.Println("c")
	default:
		fmt.Println("other")
	}

}
