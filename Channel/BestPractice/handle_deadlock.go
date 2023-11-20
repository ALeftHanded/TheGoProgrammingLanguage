package BestPractice

func SingleGoroutineReceiveData() {
	ch := make(chan int)
	ch <- 1
}

func CircleWait() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		<-ch1
		ch2 <- 1
	}()

	go func() {
		<-ch2
		ch1 <- 1
	}()

	select {} // 死锁：两个goroutine都在等待对方发送数据
}

func CloseReceiveData() {
	ch := make(chan int)
	close(ch)
	ch <- 1
}
