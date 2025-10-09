package main

import (
	"fmt"
	"time"
)

func stu21() {
	ch := make(chan int, 3)
	go NumProducter(&ch)
	go NumConsumer(ch)

	time.Sleep(3 * time.Second)
}

// 不推荐chan指针的使用，英文指针可能指向其他chan
func NumProducter(ch *chan int) {
	for i := 1; i <= 10; i++ {
		*ch <- i
	}
}

func NumConsumer(ch chan int) {
	for {
		num := <-ch
		fmt.Printf("num = %d\n", num)
	}
}
