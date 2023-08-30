package main

import (
	"image"
	"image/png"
	"math/rand"
	"os"

	"github.com/google/uuid"
	"github.com/kbinani/screenshot"
)

func TakeScreenshot() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		ScreenShotStack = append(ScreenShotStack, img)
	}
}

func ImageToFile(img *image.RGBA) (filename string) {
	id := uuid.New()
	filename = id.String() + ".png"
	file, _ := os.Create(filename)
	defer file.Close()
	png.Encode(file, img)
	return
}

func GetRandomScreensot(ScreenShootStack []*image.RGBA) (img *image.RGBA) {
	randomIndex := rand.Intn(len(ScreenShootStack))
	img = ScreenShootStack[randomIndex]
	return
}
