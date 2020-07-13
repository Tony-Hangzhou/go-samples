package grammar

import (
	"fmt"
	"time"
)

func task(id int) {

	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}

}

func Goroutine() {
	go task(1) //goroutine 创建一个并发任务单元，放置在系统队列中，等待调度
	go task(2)

	time.Sleep(time.Second * 6)
}
