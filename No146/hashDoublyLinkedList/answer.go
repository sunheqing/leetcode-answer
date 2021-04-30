package main

import "fmt"

// LRU缓存机制

// 双向链表节点
type Node struct {
	value    int
	previous *Node
	next     *Node
}

// 哈希表加双向链表
type LRUCache struct {
	capacity int           // 容量
	head     *Node         // 双向链表头节点指针
	kv       map[int]*Node // 存放缓存信息
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, kv: map[int]*Node{}, head: nil}
}

func (this *LRUCache) Get(key int) int {
	if node, have := this.kv[key]; have {
		if node != this.head {
			// 将双向链表中该节点从当前位置移动到链表头
			if node.next == nil { // 节点在链表末尾时
				node.previous.next = nil
			} else {
				node.previous.next = node.next
				node.next.previous = node.previous
			}
			node.previous = nil
			node.next = this.head
			this.head.previous = node
			// 更新头节点指针
			this.head = node
		}
		return this.head.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.capacity <= 0 { // 当容量小于等于0时不允许写入缓存
		return
	}
	if node, have := this.kv[key]; have { // 更新操作
		if node != this.head {
			// 将双向链表中该节点从当前位置移动到链表头
			if node.next == nil { // 节点在链表末尾时
				node.previous.next = nil
			} else {
				node.previous.next = node.next
				node.next.previous = node.previous
			}
			node.previous = nil
			node.next = this.head
			this.head.previous = node
			// 更新头节点指针
			this.head = node
		}
		this.head.value = value
	} else { // 新增操作
		if len(this.kv) == this.capacity { // 需要删除最久未使用的节点
			node := this.head
			if this.capacity > 1 { // 当容量大于1时再遍历链表，等于1时直接删除链表唯一节点以及键值对
				for node.next != nil {
					node = node.next
				}
				node.previous.next = nil
				node.previous = nil
			}
			for k, n := range this.kv {
				if n == node {
					delete(this.kv, k)
					break
				}
			}
		}
		this.kv[key] = &Node{value: value, previous: nil, next: nil}
		if this.head != nil {
			this.head.previous = this.kv[key]
			this.kv[key].next = this.head
		}
		// 更新头节点指针
		this.head = this.kv[key]
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
	cacheObj.Put(2, 1)
	cacheObj.Put(1, 1)
	cacheObj.Put(2, 3)
	cacheObj.Put(4, 1)
	fmt.Println(cacheObj.Get(1))
	fmt.Println(cacheObj.Get(2))
}
