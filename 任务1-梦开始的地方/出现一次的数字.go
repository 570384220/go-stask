package main

import (
	"encoding/json"
	"fmt"
)

func stu13() {

	res := onceNum()
	fmt.Println(res)

}

func onceNum() (res int) {
	arr := [...]int{1, 2, 3, 4, 5, 1, 2, 4, 5, 3}
	arrMap := map[int]int{}

	for _, v := range arr {
		var count = arrMap[v]
		count++
		arrMap[v] = count
	}

	b, _ := json.Marshal(arrMap)
	fmt.Println(string(b))
	for k, v := range arrMap {
		if v == 1 {
			return k
		}
	}

	return res
}
