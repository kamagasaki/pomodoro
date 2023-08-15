package main

import (
	"fmt"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func simpleCountdown(target time.Time, formatter func(time.Duration) string) {
	for range time.Tick(100 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Print("Countdown: ", formatter(0), "   \r")
			return
		}
		fmt.Fprint(os.Stdout, "Countdown: ", formatter(timeLeft), "   \r")
		os.Stdout.Sync()
	}
}

func TakeScreenshot() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}
