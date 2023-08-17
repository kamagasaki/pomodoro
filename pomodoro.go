// pormodoro timer
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
)

func main() {
	nokiaTune()
	start := time.Now()
	finish, err := waitDuration(start)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)
	fmt.Printf("Start timer for %s.\n\n", wait)
	formatter := formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}
	beeep.Notify("Pomodoro Info", "Start Melakukan Task 25 menit", "assets/information.png")
	X, Y := robotgo.GetMousePos()
	fmt.Println(X, Y)
	simpleCountdown(finish, formatter)
	TakeScreenshot()

	nokiaTune()
	start = time.Now()
	finish, err = waitDuration(start)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait = finish.Sub(start)
	fmt.Printf("Start timer for %s.\n\n", wait)
	formatter = formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}
	beeep.Notify("Pomodoro Info", "STOP!!!! Break Dulu 5 menit", "assets/information.png")
	TakeScreenshot()
	fmt.Println(X, Y)
	simpleCountdownBreak(finish, formatter, X, Y)

}
