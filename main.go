package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	maxTokens int
	capacity  int
	fillRate  int
}

func (t *TokenBucket) fill() {
	for {
		if t.capacity < t.maxTokens {
			fmt.Printf("[%s]fill bucket: %d\n", time.Now().Format(time.RFC3339), t.fillRate)
			t.capacity += t.fillRate
		} else {
			fmt.Printf("[%s]bucket is full\n", time.Now().Format(time.RFC3339))
		}
		time.Sleep(1 * time.Second)
	}
}

func consume(t *TokenBucket, id int) {
	for {
		if t.capacity > 0 {
			fmt.Printf("[%s][%d]consume bucket: %d\n", time.Now().Format(time.RFC3339), id, t.capacity)
			t.capacity -= 1
		} else {
			fmt.Printf("[%s][%d]bucket is empty\n", time.Now().Format(time.RFC3339), id)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	bucket := &TokenBucket{
		maxTokens: 100,
		capacity:  30,
		fillRate:  6,
	}

	done := make(chan bool)
	go bucket.fill()
	go consume(bucket, 1)
	go consume(bucket, 2)
	go consume(bucket, 3)
	<-done
}
