package grammar

func Slice() {
	x := make([]int, 0, 5) // 动态数组功能

	for i := 0; i < 8; i++ {
		x = append(x, i)
	}

	println(x)
}
