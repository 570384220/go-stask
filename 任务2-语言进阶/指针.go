package main

import "fmt"

func stu16() {
	num := 10
	add10(&num)
	fmt.Println(num)

	arr := []int{1, 2, 4}
	mutiplyTwo(&arr)
	fmt.Println(arr)

	m := map[string]int{}
	m["a"] = 1
	mapTest(m)
	fmt.Println(m)

	arrb := []int{1, 2, 4}
	mutiplyTwo(&arrb)

}

func add10(num *int) {
	*num += 10
}

// 切片本身就是一个引用了，很少使用切片指针
func mutiplyTwo(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
}

func mapTest(m map[string]int) {
	m["a"] = 2
}
