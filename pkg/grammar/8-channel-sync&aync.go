package grammar

/**
不带缓存通道，发送和接受方是阻塞的
带缓存通道，发送和接受方式异步的
*/
func ChannelSync() {

	a, b := make(chan int) /**不带缓存的通道**/, make(chan int, 3) /**带缓存通道**/

	b <- 2
	b <- 3

	println("a:", len(a), cap(a)) // 根据len和cap判断是否是阻塞chan
	println("b:", len(b), cap(b))
}
