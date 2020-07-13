package grammar

import "fmt"

func For() {

	data := []int{1, 2, 3}
	for i, d := range data { // range复制了一份data，用于迭代
		data[i] += 100
		fmt.Printf("d:%d data:%d \n", d, data[i])
	}

	for _, d := range data { // 忽略 索引值 i
		fmt.Println(d)
	}

	/**
	break; continue ；goto 都可以配合 Label，在多层嵌套中制定目标层级
	goto 不能跳转到其他函数和与自己比内层的代码块
	*/
outer: // 标签 Label
	for x := 0; x < 5; x++ {
		for y := 0; y < 10; y++ {
			if y > 2 {
				continue outer
			}

			if x > 2 {
				break outer
			}

			fmt.Printf("x:%d y:%d \n", x, y)
		}
	}
}
