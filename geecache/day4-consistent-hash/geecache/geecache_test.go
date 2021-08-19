package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	// GetterFunc类型实现了Getter接口的Get方法。
	var f Getter = GetterFunc(func(s string) ([]byte, error) {
		return []byte(s), nil
	})

	expect := []byte("key1")
	if v, _ := f.Get("key1"); !reflect.DeepEqual(expect, v) {
		t.Errorf("call back error!")
	}
}

// ---- 测试Get------------------------------------------------------------

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int)
	gee := NewGroup("scores", 2<<10, GetterFunc(
		//func(string) ([]byte, error)
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				}
				loadCounts[key] += 1
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	for k, v := range db {
		if view, ok := gee.Get(k); ok != nil || view.String() != v {
			t.Fatal("failed to get value of Tom")
		} // load from callback function
		if _, err := gee.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s miss", k)
		} // cache hit
	}

	if view, err := gee.Get("unknown"); err == nil {
		t.Fatalf("the value of unknow should be empty, but %s got", view)
	}
}
