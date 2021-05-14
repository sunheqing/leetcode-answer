package main

import "fmt"

// 最小栈
type element struct {
	val int
}
type MinStack struct {
	min   *element
	stack []*element
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: nil, stack: []*element{}}
}

func (this *MinStack) Push(val int) {
	e := &element{val: val}
	this.stack = append(this.stack, e)
	if this.min == nil || val < this.min.val {
		this.min = e
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	e := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) == 0 {
		this.min = nil
		return
	}
	if this.min == e {
		this.min = this.stack[0]
		for _, elmt := range this.stack {
			if elmt.val < this.min.val {
				this.min = elmt
			}
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	if this.min == nil {
		return 0
	}
	return this.min.val
}

func main() {
	ms := Constructor()
	ms.Push(-2)
	ms.Push(0)
	ms.Push(-3)
	fmt.Println(ms.GetMin())
	ms.Pop()
	fmt.Println(ms.Top())
	fmt.Println(ms.GetMin())
}
