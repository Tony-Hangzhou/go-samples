package grammar

import (
	"errors"
	"fmt"
	"log"
)

func check(i int) error {
	if i == 0 {
		errors.New("i 不能小于 0")
	}
	return nil
}

func If() {

	i := 0

	if e := check(i); e != nil { // 优先执行 check 函数
		log.Fatalln(e)
	} else {
		i++
		fmt.Println(i)
	}
}
