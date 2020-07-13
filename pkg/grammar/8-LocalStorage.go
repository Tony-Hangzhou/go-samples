package grammar

import (
	"fmt"
	"sync"
)

/**
与线程不同，gorutine 无法设置优先级、编号、无局部存储、返回值被抛弃
*/
func LocalStorage() {

	var wg sync.WaitGroup
	// 定义结构体数组，为 gorutine 设置编号和保存返回结果
	var gs [5]struct {
		id     int
		result int
	}

	for i := 0; i < len(gs); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			gs[id].id = id
			gs[id].result = (id + 1) * 100

		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v\n", gs)

}
