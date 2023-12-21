package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	// Сценарий 1: один из каналов закрывается
	t.Run("Closes when one channel is closed", func(t *testing.T) {
		start := time.Now()
		c1 := make(chan interface{})
		c2 := make(chan interface{})

		go func() {
			time.Sleep(50 * time.Millisecond)
			close(c1)
		}()

		select {
		case <-or(c1, c2):
		case <-time.After(100 * time.Millisecond):
			t.Fatal("Expected or to close")
		}

		if time.Since(start) > 100*time.Millisecond {
			t.Error("Expected to close within 100ms")
		}
	})

	// Сценарий 2: нет входных каналов
	t.Run("Handles no channels", func(t *testing.T) {
		select {
		case <-or():
		case <-time.After(50 * time.Millisecond):
			t.Fatal("Expected or to close immediately when no channels are passed")
		}
	})

}
