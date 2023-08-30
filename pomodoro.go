// pormodoro timer
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Printf("HostName is: %s\n", Hostname)
	WhatsApp()
	//SendReportTo(filename, "6281313112053-1492882006")

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

	fmt.Println("1 sesi pomodoro selesai")
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	WAclient.Disconnect()

}
