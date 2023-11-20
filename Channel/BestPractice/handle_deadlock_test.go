package BestPractice

import "testing"

func TestDeadLockExample(t *testing.T) {
	// fatal error: all goroutines are asleep - deadlock!
	t.Run("SingleGoroutineReceiveData", func(t *testing.T) {
		SingleGoroutineReceiveData()
	})
	// fatal error: all goroutines are asleep - deadlock!
	t.Run("CircleWait", func(t *testing.T) {
		CircleWait()
	})
	// panic: send on closed channel
	t.Run("CloseReceiveData", func(t *testing.T) {
		CloseReceiveData()
	})
}
