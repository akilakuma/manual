package main

import "sync"

// ❌ 錯誤示例：Pool 中存放有狀態的物件
type BadBuffer struct {
	data []byte
	used bool // 這個狀態可能導致問題
}

var badPool = sync.Pool{
	New: func() interface{} {
		return &BadBuffer{
			data: make([]byte, 0, 1024),
			used: false,
		}
	},
}

func useBadPool() {
	buf := badPool.Get().(*BadBuffer)
	buf.used = true // 修改狀態

	// 如果忘記重置狀態就放回 pool，下次使用會有問題
	badPool.Put(buf)
}
