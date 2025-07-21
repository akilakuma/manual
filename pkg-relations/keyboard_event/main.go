package main

import (
	kb "github.com/micmonay/keybd_event"
	"time"
)

func main() {
	// 建立 keybd 物件
	k, err := kb.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// 指定要送出的按鍵
	k.SetKeys(kb.VK_SPACE)

	// 移除修飾鍵（避免誤觸 Ctrl+Shift 等）
	k.HasCTRL(false)
	k.HasALT(false)

	for {
		// 模擬按下
		k.Launching()
		time.Sleep(100 * time.Millisecond)
		// 模擬放開（這其實可省略，但完整起見）
		k.Release()

		// 間隔時間 println("Fake key sent to prevent screensaver.")
		time.Sleep(150 * time.Second)
	}
}
