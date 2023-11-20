package main

import (
	"fmt"
	"sync"
	"time"
)

func sendInts(nums []int, ch chan<- int) {
	// 题目1：单向通道
	//请实现一个函数，它接受一个整数切片和一个单向通道，将切片中的所有整数发送到通道。
	defer close(ch) // 确保发送完所有整数后关闭通道
	for _, num := range nums {
		ch <- num
	}
}
func receiveInts(ch <-chan int) []int {
	// 题目2：通道关闭和检测
	//请实现一个函数，它接受一个整数通道和一个整数切片。它应该从通道中读取整数，直到通道关闭，
	// 然后将所有读取到的整数存储到切片中。请确保你的实现能够正确处理通道关闭的情况。
	var res []int
	for num := range ch {
		res = append(res, num)
	}
	return res
}

func syncC() {
	//题目3：通道同步
	//请使用两个goroutine和一个通道来实现同步。第一个goroutine应该打印从1到5的整数，
	//每打印一个数字后，它应该等待第二个goroutine的信号才能继续打印。第二个goroutine应该负责发送信号。
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			<-ch
			println(i)
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- struct{}{}

		}
	}()

	wg.Wait()
}

func workers4(nums []int) {
	//题目4：带缓冲的通道和工作池
	//实现一个工作池，其中有4个工作goroutine来处理整数切片中的整数。
	//每个工作goroutine应该从带缓冲的通道接收整数，对其进行平方处理，然后打印结果。
	//主goroutine应该负责将整数发送到通道，并在所有工作完成后关闭通道。
	ch := make(chan int, len(nums))
	for _, num := range nums {
		ch <- num
	}
	close(ch)

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for tmp := range ch {
				fmt.Println(tmp * tmp)
				time.Sleep(500 * time.Millisecond)
			}

		}()
	}

	wg.Wait()
}

func main() {
	//ch := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	sendInts([]int{1, 2, 3, 4, 5}, ch)
	//
	//}()
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	fmt.Println(receiveInts(ch))
	//}()
	//wg.Wait()

	//syncC()

	workers4([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
