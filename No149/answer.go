package main

import "fmt"

type mark struct {
	open bool
	num  int
}

// todo 直线上最多的点
func maxPoints(points [][]int) int {
	length := len(points)
	if length <= 2 {
		return length
	}
	var angle interface{}
	m := make(map[interface{}]*mark)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if points[i][1] == points[j][1] {
				angle = ""
			} else {
				angle = float64(points[i][0]-points[j][0]) / float64(points[i][1]-points[j][1])
			}
			if _, have := m[angle]; have {
				if m[angle].open {
					m[angle].num += 1
				}
			} else {
				m[angle] = &mark{open: true, num: 2}
			}
		}
		for _, k := range m {
			k.open = false
		}
	}
	max := 0
	for k, v := range m {
		fmt.Println(k, *v)
		if v.num > max {
			max = v.num
		}
	}
	return max
}

func main() {
	points := [][]int{[]int{1, 1}, []int{3, 2}, []int{5, 3}, []int{4, 1}, []int{2, 3}, []int{1, 4}}
	points2 := [][]int{[]int{1, 1}, []int{2, 2}, []int{3, 3}}
	points3 := [][]int{[]int{2, 3}, []int{3, 3}, []int{-5, 3}}
	points4 := [][]int{[]int{0, 0}, []int{4, 5}, []int{7, 8}, []int{8, 9}, []int{5, 6}, []int{3, 4}, []int{1, 1}}
	fmt.Println(maxPoints(points))
	fmt.Println("--------------------")
	fmt.Println(maxPoints(points2))
	fmt.Println("--------------------")
	fmt.Println(maxPoints(points3))
	fmt.Println("--------------------")
	fmt.Println(maxPoints(points4))
}
