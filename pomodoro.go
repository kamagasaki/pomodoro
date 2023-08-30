// pormodoro timer
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/gen2brain/beeep"
)

func main() {
	wag := os.Getenv("POMOGROUPWA")
	if wag == "" {
		go DownloadFile("information.png", "https://pomokit.github.io/pomodoro/information.png")
		go DownloadFile("warning.png", "https://pomokit.github.io/pomodoro/warning.png")
		beeep.Alert("Pomokit Info", "Please Input Your WhatsApp Group ID from Iteung", "information.png")
		fmt.Println("Please Input Your WhatsApp Group ID from Iteung : ")
		fmt.Scanln(&wag)
		wag = strings.TrimSpace(wag)
		os.Setenv("POMOGROUPWA", wag)
	}

	WhatsApp()

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("task")
	GetSetTime("break")

	GetSetTime("longbreak")

	img := GetRandomScreensot(ScreenShotStack)
	filename := ImageToFile(img)
	SendReportTo(filename, wag)
	msg := "Selamat!!!!! 1 sesi pomodoro selesai dengan jumlah skrinsutan:" + strconv.Itoa(len(ScreenShotStack))

	beeep.Alert("Pomokit Info", msg, "information.png")
	fmt.Println("1 sesi pomodoro selesai dengan jumlah skrinsutan:", strconv.Itoa(len(ScreenShotStack)))
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	WAclient.Disconnect()

}
