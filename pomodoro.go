// pormodoro timer
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gen2brain/beeep"
)

func main() {

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
	SendReportTo(filename, "6281313112053-1492882006")
	msg := "Selamat!!!!! 1 sesi pomodoro selesai dengan jumlah skrinsutan:" + strconv.Itoa(len(ScreenShotStack))

	beeep.Alert("Pomodoro Info", msg, "assets/information.png")
	fmt.Println("1 sesi pomodoro selesai dengan jumlah skrinsutan:", strconv.Itoa(len(ScreenShotStack)))
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	WAclient.Disconnect()

}
