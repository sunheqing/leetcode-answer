package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 验证回文串，给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i <= j {
		if unicode.IsLetter(rune(s[i])) == false && unicode.IsDigit(rune(s[i])) == false {
			i += 1
		} else if unicode.IsLetter(rune(s[j])) == false && unicode.IsDigit(rune(s[j])) == false {
			j -= 1
		} else if s[i] == s[j] {
			i += 1
			j -= 1
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
