package geecache

import (
	"fmt"
	"log"
	"sync"
)

/**
GetterFunc 实现了Get方法。
Getter接口类型 可以被 GetterFunc类型的值赋值
*/

type GetterFunc func(string) ([]byte, error)

type Getter interface {
	Get(string) ([]byte, error)
}

/**
函数类型实现某一个接口，称之为接口型函数，
方便使用者在调用时既能够传入函数作为参数，也能够传入实现了该接口的结构体作为参数。
把函数转换为接口的函数
*/

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// ---- Group --------------------------------------------------------------------------

type Group struct {
	name      string
	getter    Getter // Getter 接口类型
	mainCache cache  // cache 结构体
}

var mu sync.RWMutex
var groups = make(map[string]*Group)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := Group{
		name:   name,
		getter: getter,
		mainCache: cache{
			cacheBytes: cacheBytes,
		},
	}
	groups[name] = &g
	return &g
}

func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	return groups[name]
}

// ---- Get --------------------------------------------------------------------------

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		// func Errorf(format string, a ...interface{}) error {....
		return ByteView{}, fmt.Errorf("key is miss!")
	}

	if v, ok := g.mainCache.get(key); ok {
		// 找到了
		log.Println("[GeeCache] hit")
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}

	value := ByteView{cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
