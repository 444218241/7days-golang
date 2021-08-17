package lru

import (
	"container/list"
)

type Value interface {
	Len() int // 用于返回值所占用的内存大小。
}

type Cache struct {
	maxBytes  int64 // 允许使用的最大内存
	nbyes     int64 // 当前已使用的内存(key和value的大小和)
	ll        *list.List
	cache     map[string]*list.Element // 键是字符串，值是双向链表中对应节点的指针。
	OnEvicted func(string, Value)      // 某条记录被移除时的回调函数，可以为 nil。
}

// 双向链表节点的数据类型
type entry struct {
	key   string
	value Value
}

// New is the Constructor of Cache
// 查找主要有 2 个步骤，第一步是从字典中找到对应的双向链表的节点，第二步，将该节点移动到队尾。
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		nbyes:     0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
// 这里的删除，实际上是缓存淘汰。即移除最近最少访问的节点（队首）
// 双向链表作为队列，队首队尾是相对的，在这里约定 front 为队尾
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbyes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增/修改
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbyes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		c.cache[key] = ele
		c.nbyes += int64(value.Len()) + int64(len(key))
	}
	if c.maxBytes != 0 && c.nbyes > c.maxBytes {
		c.RemoveOldest()
	}
}
