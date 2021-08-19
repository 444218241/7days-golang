package consistenthash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	// return *Map
	// 这里我们自定义了计算hash的方法
	m := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	// 2, 4, 6, 12, 14, 16, 22, 24, 26
	m.Add("6", "4", "2")

	// key是虚拟节点，value是真实的节点
	testCases := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}
	fmt.Printf("m first %+v \n", m.hashMap)
	// [2:2 4:4 6:6 12:2 14:4 16:6 22:2 24:4 26:6]
	for k, v := range testCases {
		if m.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %q, But Get() is %s", k, v, m.Get(k))
		}
	}

	m.Add("8")
	fmt.Printf("m after add new key %+v \n", m.hashMap)
	// [2:2 4:4 6:6 8:8 12:2 14:4 16:6 18:8 22:2 24:4 26:6 28:8]
	// 发现，没有发生雪崩，也就增加了 8，18，28对应的value都是8
	// 27 should now map to 8.
	testCases["27"] = "8"
	for k, v := range testCases {
		if m.Get(k) != v {
			t.Errorf("Asking for %s, should have yielded %s, But Get() is %s", k, v, m.Get(k))
		}
	}
}
