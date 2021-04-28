package main

import "fmt"

// 只出现一次的数字，给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
func singleNumber(nums []int) int {
	record := make(map[int]int)
	for _, num := range nums {
		if _, have := record[num]; have {
			delete(record, num)
		} else {
			record[num] = 0
		}
	}
	for k := range record {
		return k
	}
	panic(nums)
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
}
