package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func stu22() {
	num := Number{number: 0}
	go add(&num)
	go add(&num)

	time.Sleep(2 * time.Second)
	fmt.Println("num1: ", num.number)

	num2 := int64(0)
	go add2(&num2)
	go add2(&num2)
	fmt.Println("num2: ", num.number)

	time.Sleep(1 * time.Second)
}

type Number struct {
	mutex  sync.Mutex
	number int
}

func add(num *Number) {
	for i := 0; i < 1000; i++ {
		num.mutex.Lock()
		num.number += 1
		num.mutex.Unlock()
	}
}
func add2(num *int64) {
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(num, 1)
	}
}
