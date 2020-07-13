package grammar

func Exp() {

	/**
	for
	*/
	for i := 0; i < 5; i++ {
		println(i)
	}

	/**
	while
	*/
	x := 0
	for x < 5 {
		println(x)
		x++
	}
	/**
	while(true)
	*/
	x += 0
	for {
		println(x)
		x++
		if x > 5 {
			break
		}
	}
}
