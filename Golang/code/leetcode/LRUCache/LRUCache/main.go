/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.

Follow up:
Could you do both operations in O(1) time complexity?

链接：https://leetcode-cn.com/problems/lru-cache
*/
package LRUCache

import "container/list"

type item struct {
	k int
	v int
}

type LRUCache struct {
	myCap int
	mp    map[int]*list.Element
	lst   *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		myCap: capacity,
		mp:    make(map[int]*list.Element),
		lst:   list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.mp[key]; !ok {
		return -1
	}
	kv := c.lst.Remove(c.mp[key]).(item)
	c.mp[key] = c.lst.PushFront(kv)
	return kv.v
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.mp[key]; ok {
		c.lst.Remove(c.mp[key])
	} else if len(c.mp) == c.myCap {
		back := c.lst.Back()
		backKV := back.Value.(item)
		c.lst.Remove(back)
		delete(c.mp, backKV.k)
	}
	c.mp[key] = c.lst.PushFront(item{k: key, v: value})
}
