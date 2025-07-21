package main

import (
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

/*
	螢幕截圖好用的package
	注意要在主作業系統內執行，wsl或docker等會有問題，因為是調用os system api
*/

func main() {

	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		fileName := "screenshot_" + string('0'+i) + ".png"
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}
}
