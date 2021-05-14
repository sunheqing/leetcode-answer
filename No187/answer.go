package main

import (
	"fmt"
	"unsafe"
)

// 重复的DNA序列
// 使用哈希表计数
func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	subCount := make(map[string]int)
	result := []string{}
	for i := 0; i <= len(s)-10; i++ {
		if _, have := subCount[s[i:i+10]]; have {
			subCount[s[i:i+10]] += 1
		} else {
			subCount[s[i:i+10]] = 1
		}
	}
	for sub, count := range subCount {
		if count > 1 {
			result = append(result, sub)
		}
	}
	return result
}

// todo 使用哈希表计数，优化版，利用A、C、G、T这个信息
func findRepeatedDnaSequences2(s string) []string {
	if len(s) <= 10 {
		return []string{}
	}
	result := []string{}
	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(unsafe.Sizeof("ACGTACGTAC"))
	fmt.Println(unsafe.Sizeof("0000000000"))
}
