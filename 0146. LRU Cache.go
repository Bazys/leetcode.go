package leetcode

import (
	"container/list"
)

type Item struct {
	Key   int
	Value int
}

type LRUCache struct {
	capacity int
	items    map[int]*list.Element
	queue    *list.List
}

func Constructor(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	element, exists := c.items[key]
	if !exists {
		return -1
	}
	c.queue.MoveToFront(element)
	return element.Value.(*Item).Value
}

func (c *LRUCache) purge() {
	if element := c.queue.Back(); element != nil {
		key := c.queue.Remove(element).(*Item).Key
		delete(c.items, key)
	}
}

func (c *LRUCache) Put(key int, value int) {
	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return
	}

	if c.queue.Len() == c.capacity {
		c.purge()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[key] = element
}
