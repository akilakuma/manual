package main

import "sync"

// ✅ 正確示例：Pool 中存放無狀態物件，或正確重置狀態
type GoodBuffer struct {
	data []byte
}

func (b *GoodBuffer) Reset() {
	b.data = b.data[:0] // 重置 slice 長度但保留 capacity
}

var goodPool = sync.Pool{
	New: func() interface{} {
		return &GoodBuffer{
			data: make([]byte, 0, 1024),
		}
	},
}

func useGoodPool() {
	buf := goodPool.Get().(*GoodBuffer)
	defer func() {
		buf.Reset() // 使用完畢後重置狀態
		goodPool.Put(buf)
	}()

	// 使用 buffer
	buf.data = append(buf.data, []byte("some data")...)
}
