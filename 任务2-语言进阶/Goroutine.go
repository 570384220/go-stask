package main

import (
	"fmt"
	"math/rand"
	"time"
)

func stu17() {

	go func() {
		arr := make([]int, 10)
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				arr = append(arr, i)
			}
		}
		fmt.Println("1到10的偶数：", arr)
	}()

	go func() {
		arr := []int{}
		for i := 0; i < 10; i++ {
			if i%2 != 0 {
				arr = append(arr, i)
			}
		}
		fmt.Println("1到10的奇数：", arr)
	}()

	time.Sleep(2 * time.Second)
}

func stu18() {
	a := 8
	fmt.Println(time.Duration(a)) // 8ns 8纳秒

	task1 := func(num int) {
		begin := time.Now()
		time.Sleep(time.Duration(num) * time.Second)
		end := time.Now()
		fmt.Println("任务1耗时：", end.Sub(begin))
	}

	task2 := func(num int) {
		begin := time.Now()
		time.Sleep(time.Duration(num) * time.Second)
		end := time.Now()
		fmt.Println("任务2耗时：", end.Sub(begin))
	}

	tasks := []func(int){task1, task2}
	taskExecutor(tasks)

	time.Sleep(10 * time.Second)
}

func taskExecutor(tasks []func(int)) {
	for _, task := range tasks {
		num := rand.Intn(10)
		go task(num)
	}
}
