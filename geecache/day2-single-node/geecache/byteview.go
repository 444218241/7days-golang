package geecache

/**
我们抽象了一个只读数据结构 ByteView 用来表示缓存值，是 GeeCache 主要的数据结构之一。
b 是只读的，使用 ByteSlice() 方法返回一个拷贝，防止缓存值被外部程序修改。
*/
type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
