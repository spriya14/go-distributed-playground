package common

type Args struct {
	A, B    int
	Payload []byte
}

type Reply struct {
	Result int
}

type PayloadSize struct {
	name string
	size int
}
