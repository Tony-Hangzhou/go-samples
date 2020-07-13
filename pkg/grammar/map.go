package grammar

func Map() {

	m := make(map[string]int)
	m["a"] = 1

	x, ok := m["b"] // ok-idiom模式，ok 值表示是否成功
	println(x, ok)

}
