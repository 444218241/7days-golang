package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to int, Hash type is a func Type
type Hash func([]byte) uint32

// Map constains all hashed keys
type Map struct {
	hash     Hash           // is not custom if will be set as "func ChecksumIEEE(data []byte) uint32 {}"
	replicas int            // 虚拟节点倍数
	keys     []int          // Sorted, Slice Of 哈希环
	hashMap  map[int]string // map: 虚拟节点与真实节点的映射表 hashMap，键是「虚拟节点的哈希值」(恰好键不能重复)，值是「真实节点的名称」。
}

// New creates a Map instance
// 构造函数 New() 允许自定义虚拟节点倍数和 Hash 函数。
func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys(真实存储节点) to the hash
func (m *Map) Add(keys ...string) {
	// 遍历传入的真实节点
	for _, key := range keys {
		// 增加虚拟节点
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash) // 增加虚拟节点
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys) // 虚拟节点组成的哈希环排序
}

// Get gets the closest  item in the hash to the provided给予 key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
