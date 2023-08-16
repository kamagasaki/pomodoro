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
	start := time.Now()

	finish, err := waitDuration(start)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	wait := finish.Sub(start)

	formatter := formatSeconds
	switch {
	case wait >= 24*time.Hour:
		formatter = formatDays
	case wait >= time.Hour:
		formatter = formatHours
	case wait >= time.Minute:
		formatter = formatMinutes
	}

	fmt.Printf("Start timer for %s.\n\n", wait)
	TakeScreenshot()
	nokiaTune()
	X, Y := robotgo.GetMousePos()
	fmt.Println(X, Y)
	err = beeep.Notify("Pomodoro Info", "Start Melakukan Task 25 menit", "assets/information.png")
	if err != nil {
		panic(err)
	}

	if *simple {
		simpleCountdown(finish, formatter)
	} else {
		fullscreenCountdown(start, finish, formatter)
	}
	/* 	err = beeep.Alert("Break", "Break Dulu Selama 5 Menit", "assets/warning.png")
	   	if err != nil {
	   		panic(err)
	   	} */

	if !*silence {
		fmt.Println("\a") // \a is the bell literal.
	} else {
		fmt.Println()
	}
}
