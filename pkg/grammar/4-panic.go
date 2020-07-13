package grammar

/**
panic recover 相当于 try catch
服务端端口被占用或者数据库未启动情况会触发
*/
import (
	"fmt"
)

func testException() {
	defer println("test.1")
	defer println("test.2")

	panic("i am dead") // 触发宕机，发生致命性错误

	fmt.Println("aaaa") // panic触发了宕机，所以该句不会被执行
}

func Panic() {
	defer func() {
		if info := recover(); info != nil { // recover 捕获到发生致命性错误，做后事处理。连续触发panic，仅最后一个会被 recover 捕获。
			fmt.Println("触发了宕机", info)
		}
	}()

	testException()
}
