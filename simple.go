package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-vgo/robotgo"
)

func simpleCountdown(target time.Time, formatter func(time.Duration) string) {
	var takescreenshoot bool
	timeLeft := -time.Since(target)
	minutetake := rand.Int63n(int64(timeLeft.Minutes()))
	for range time.Tick(100 * time.Millisecond) {
		timeLeft = -time.Since(target)
		if timeLeft < 0 {
			fmt.Print("Countdown: ", formatter(0), "   \r")
			return
		}
		if int64(timeLeft.Minutes()) == minutetake && !takescreenshoot {
			TakeScreenshot()
			takescreenshoot = true
		}
		fmt.Fprint(os.Stdout, "Countdown: ", formatter(timeLeft), "   \r")
		os.Stdout.Sync()
	}
}

func SimpleBreakCountdown(target time.Time, formatter func(time.Duration) string, X, Y int) {
	for range time.Tick(100 * time.Millisecond) {
		timeLeft := -time.Since(target)
		if timeLeft < 0 {
			fmt.Print("Countdown: ", formatter(0), "   \r")
			return
		}
		robotgo.DragSmooth(X, Y)
		fmt.Fprint(os.Stdout, "Countdown: ", formatter(timeLeft), "   \r")
		os.Stdout.Sync()
	}
}

func GetSetTime(status string) (finish time.Time, formatter func(time.Duration) string, err error) {
	nokiaTune()
	start := time.Now()
	if status == "task" {
		finish, err = waitDuration(start)
	} else if status == "break" {
		finish, err = waitBreakDuration(start)
	} else {
		finish, err = waitDuration(start)
	}

	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)
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
	if status == "task" {
		beeep.Notify("Pomodoro Info", "Start Melakukan Task 25 menit", "assets/information.png")
		simpleCountdown(finish, formatter)
	} else if status == "break" {
		beeep.Alert("Pomodoro Info", "STOP!!!! Break Dulu 5 menit", "assets/warning.png")
		X, Y := robotgo.GetMousePos()
		SimpleBreakCountdown(finish, formatter, X, Y)
	} else {
		beeep.Alert("Pomodoro Info", "STOP!!!! Istirahat Panjang Dulu 25 menit", "assets/warning.png")
		X, Y := robotgo.GetMousePos()
		SimpleBreakCountdown(finish, formatter, X, Y)
	}
	return
}
