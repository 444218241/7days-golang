package main

import (
	"container/list"
	"fmt"
)

// Value类型的值，只能由实现了Len方法的类型的值来赋值
type Value interface {
	Len() int
}

type String string

// String实现了Len方法，则String类型的值可以赋值给Value变量，且有Len方法
func (s String) Len() int {
	return len(s)
}

type Cache struct {
	ll        *list.List               // type List struct {}
	cache     map[string]*list.Element // type Element struct {}；键是字符串，值是双向链表中对应节点的指针。
	OnEvicted func(string, Value)      // type Value interface {}
}

// 键值对 entry 是双向链表节点的数据类型，
// 在链表中仍保存每个值对应的 key 的好处在于，
// 淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value // type Value interface {}
}

func main() {
	cache1 := New(func(key string, value Value) {
		fmt.Printf("key:%s, value:%s, lenOfValue:%d \n", key, value, value.Len())
	})
	// Add() 需要Value类型，而String类型实现了Len方法，所以可以作为Value类型的值传递。
	cache1.Add("china", String("BJ"))
	v, _ := cache1.Get("china")
	fmt.Println(v)
}

func New(OnEvicted func(string, Value)) *Cache {
	return &Cache{
		ll:        list.New(), // func New() *List { return new(List).Init() }
		cache:     make(map[string]*list.Element),
		OnEvicted: OnEvicted,
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)

		kv := ele.Value.(*entry)
		// Element结构体中value字段类型是interface{}，保存是时候，进行了类型断言转换
		kv.value = value
	} else {
		// func (l *List) PushFront(v interface{}) *Element
		newEntry := &entry{
			key:   key,
			value: value,
		}
		ele := c.ll.PushFront(newEntry)
		c.cache[key] = ele
	}
}

func (c Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		fmt.Printf("ele.Value v：%+v, t：%t, T：%T \n", ele.Value, ele.Value, ele.Value)
		// ele.Value v：&{key:china value:BJ}, t：&{%!t(string=china) %!t(main.String=BJ)}, T：*main.entry
		kv := ele.Value.(*entry) // 双向链表节点的数据值（类型是entry）
		/**
		Element结构体中value字段类型是interface{}，
		即需要自定义节点值的类型。
		此处把这个字段定义为了entry，
		所以要把Element中的值取出来需要用(*entry)转换。
		*/
		fmt.Printf("kv v：%+v, t：%t, T：%T \n", kv, kv, kv)
		// kv v：&{key:china value:BJ}, t：&{%!t(string=china) %!t(main.String=BJ)}, T：*main.entry
		return kv.value, true
	}
	return
}
