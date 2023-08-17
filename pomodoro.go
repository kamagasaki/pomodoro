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
	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("longbreak")

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
		TakeScreenshot()
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
