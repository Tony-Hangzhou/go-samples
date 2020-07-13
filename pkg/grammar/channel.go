package grammar

func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x)
	}

	done <- true

}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i //箭头的指向就是数据的流向
	}
	close(data)
}

func Channel() {
	done := make(chan bool)
	data := make(chan int)
	go consumer(data, done)
	go producer(data)

	<-done
}
