package grammar

/**
用 channel 在主线程阻塞，直到 gorutine 运行完成。
*/
func Channel1() {

	done := make(chan struct{})
	c := make(chan string) // 定义传输通道的数据类型

	go func() {
		s := <-c // 接受管道内容
		println(s)
		close(done)
	}()

	c <- "hi!" // 发送管道内容
	<-done     // 阻塞，直到done关闭管道
}
