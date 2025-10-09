package main

import (
	"fmt"
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
