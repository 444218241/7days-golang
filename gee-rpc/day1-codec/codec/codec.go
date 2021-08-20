package codec

type Header struct {
	ServiceMethod string // 服务名和方法名
	Seq           uint64 // 请求的序号
	error         string
}
