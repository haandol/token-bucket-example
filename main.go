package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	maxTokens int
	fillRate  int
	ch        chan int
}

func NewTokenBucket(maxTokens, fillRate int) *TokenBucket {
	return &TokenBucket{
		maxTokens: maxTokens,
		fillRate:  fillRate,
		ch:        make(chan int, maxTokens),
	}
}

func (t *TokenBucket) init(n int) {
	for i := 0; i < n; i++ {
		t.ch <- i
	}
}

func (t *TokenBucket) fill() {
	for {
		fmt.Printf("[%s]fill bucket: %d\n", time.Now().Format(time.RFC3339), t.fillRate)
		for i := 0; i < t.fillRate && len(t.ch) < t.maxTokens; i++ {
			t.ch <- i
		}
		time.Sleep(1 * time.Second)
	}
}

func consume(t *TokenBucket, id int) {
	for {
		select {
		case i := <-t.ch:
			fmt.Printf("[%s][%d]consume bucket: %d\n", time.Now().Format(time.RFC3339), id, i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	bucket := NewTokenBucket(100, 6)
	bucket.init(30)

	done := make(chan bool)
	go bucket.fill()
	go consume(bucket, 1)
	go consume(bucket, 2)
	go consume(bucket, 3)
	<-done
}
