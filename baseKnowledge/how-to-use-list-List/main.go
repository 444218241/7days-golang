package main

import (
	"container/list"
	"fmt"
)

type String string

func (s String) Len() int {
	return len(s)
}

type Cache struct {
	ll        *list.List               // type List struct {}
	cache     map[string]*list.Element // type Element struct {}
	OnEvicted func(string, Value)
}

type Value interface {
	Len() int
}

type entry struct {
	key   string
	value Value // type Value interface {}
}

func main() {
	cache1 := New(func(key string, value Value) {
		fmt.Printf("key:%s, value:%s, lenOfValue:%d \n", key, value, value.Len())
	})
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
		kv := ele.Value.(*entry)
		fmt.Printf("kv v：%+v, t：%t, T：%T \n", kv, kv, kv)
		// kv v：&{key:china value:BJ}, t：&{%!t(string=china) %!t(main.String=BJ)}, T：*main.entry
		return kv.value, true
	}
	return
}
