package main

import (
	"fmt"
	"strings"
)

// 71. 简化路径

func simplifyPath(path string) string {
	var stack []string
	for _, s := range strings.Split(path, "/") {
		if s != "" && s != "." {
			if s == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, s)
			}
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("//a/.."))
}
