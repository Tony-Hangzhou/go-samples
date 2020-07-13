package grammar

import (
	"sync"
	"time"
)

func Waitgroup() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Second)
			println("goroutine", id, "done.")

		}(i)
	}

	println("main ...")
	wg.Wait() // 等待，直到计数器为零
	println("main exit")
}
