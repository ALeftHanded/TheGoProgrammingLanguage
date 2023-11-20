package main

import (
	"fmt"
	"sync"
	"time"
)

// 请编写一个Go程序，启动两个goroutine
// 一个goroutine生成0到9的数字发送到channel
// 另一个goroutine从channel中读取这些数字并打印

// recvAndPrint 原始版本
func recvAndPrint() {
	chSend := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chSend <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			p := <-chSend
			fmt.Println(p)
		}

	}()

	time.Sleep(2 * time.Second)
}

// recvAndPrintOp 优化版本
// 1. Channel Closure: It is a good practice to close a channel when no more values will be sent on it.
// This allows the receiving goroutine to stop receiving values when the channel is closed.
// You can use the close function to close the channel after sending all numbers.
//
// 2. Wait for Completion: Instead of using time.Sleep to wait for goroutines to finish,
// which is generally not recommended due to its non-deterministic nature, you should use a sync.WaitGroup.
// This provides a more reliable way to wait for goroutines to complete their execution.
//
// 3. Buffered Channel: Since you know in advance how many numbers will be sent to the channel, you can make it a buffered channel.
// This helps to avoid the sending goroutine from being blocked until the receiving goroutine is ready to receive the value.
func recvAndPrintOp() {
	chSend := make(chan int, 10) // Buffered channel
	var wg sync.WaitGroup
	wg.Add(2) // Add 2 for two goroutines

	// Goroutine for sending numbers
	go func() {
		defer wg.Done() // Decrements counter when goroutine completes
		for i := 0; i < 10; i++ {
			chSend <- i
		}
		close(chSend) // Close the channel after sending all numbers
	}()

	// Goroutine for receiving and printing numbers
	go func() {
		defer wg.Done() // Decrements counter when goroutine completes
		for p := range chSend {
			fmt.Println(p)
		}
	}()

	wg.Wait() // Wait for both goroutines to complete
}
