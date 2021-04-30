package main

import "fmt"

// LRU缓存机制

// 带有计数器的节点
type Node struct {
	value   int
	counter int
}

// 哈希表加计数器
type LRUCache struct {
	capacity   int           // 容量
	maxCounter int           // 整个缓存中最高的计数器值
	kv         map[int]*Node // 存放缓存信息
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, maxCounter: 0, kv: map[int]*Node{}}
}

func (this *LRUCache) Get(key int) int {
	if node, have := this.kv[key]; have {
		node.counter = this.maxCounter + 1
		this.maxCounter = node.counter
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 { // 当容量小于等于0时不允许写入缓存
		return
	}
	if node, have := this.kv[key]; have { // 更新操作
		node.counter = this.maxCounter + 1
		this.maxCounter = node.counter
		node.value = value
	} else { // 新增操作
		if len(this.kv) == this.capacity { // 需要删除最久未使用的节点
			var longestUnusedKey, minCounter int
			flag := true
			for k, n := range this.kv {
				if flag || n.counter < minCounter {
					longestUnusedKey = k
					minCounter = n.counter
					flag = false
				}
			}
			delete(this.kv, longestUnusedKey)
		}
		this.kv[key] = &Node{value: value, counter: this.maxCounter + 1}
		this.maxCounter = this.kv[key].counter
	}
}

func main() {
	cacheObj := Constructor(2)
	// case1
	//fmt.Println(cacheObj.Get(2))
	//cacheObj.Put(2, 6)
	//fmt.Println(cacheObj.Get(1))
	//cacheObj.Put(1, 5)
	//cacheObj.Put(1, 2)
	//fmt.Println(cacheObj.Get(1))
	//fmt.Println(cacheObj.Get(2))

	// case2
	//cacheObj.Put(2, 1)
	//cacheObj.Put(1, 1)
	//cacheObj.Put(2, 3)
	//cacheObj.Put(4, 1)
	//fmt.Println(cacheObj.Get(1))
	//fmt.Println(cacheObj.Get(2))

	// case3
	cacheObj.Put(1, 1)
	cacheObj.Put(2, 2)
	fmt.Println(cacheObj.Get(1))
	cacheObj.Put(3, 3)
	fmt.Println(cacheObj.Get(2))
	cacheObj.Put(4, 4)
	fmt.Println(cacheObj.Get(1))
	fmt.Println(cacheObj.Get(3))
	fmt.Println(cacheObj.Get(4))
}
