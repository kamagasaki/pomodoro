package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/gen2brain/beeep"
	"github.com/whatsauth/watoken"
)

func main() {
	if !FileExists("information.png") {
		go DownloadFile("information.png", InfoImageURL)
	}
	if !FileExists("warning.png") {
		go DownloadFile("warning.png", WarningImageURL)
	}

	wag := FiletoString("wag.info")
	if wag == "" {
		beeep.Alert("Pomokit Info", "Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ", "information.png")
		fmt.Println("Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ")
		fmt.Scanln(&wag)
		wag = strings.TrimSpace(wag)
		wag = strings.ReplaceAll(wag, " ", "")
		StringtoFile(wag, "wag.info")
		os.Setenv("POMOGROUPWA", wag)
	}
	urltask := FiletoString("id.info")
	if urltask == "" {
		fmt.Println("URL Github Pages Yang Akan Dikerjakan : ")
		fmt.Scanln(&urltask)
		urltask = strings.TrimSpace(urltask)
		urltask = strings.ReplaceAll(urltask, " ", "")
		StringtoFile(urltask, "id.info")
		os.Setenv("USERIDPOMO", urltask)
	}
	hashuserid, err := watoken.EncodeforHours(urltask, PrivateKey, 3)
	if err != nil {
		fmt.Println(err)
	}

	WhatsApp()
	SendNotifTo(wag)

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
	SendReportTo(filename, wag, hashuserid)
	msg := "Selamat!!!!! 1 sesi pomodoro selesai dengan jumlah skrinsutan:" + strconv.Itoa(len(ScreenShotStack))

	beeep.Alert("Pomokit Info", msg, "information.png")
	fmt.Println("1 sesi pomodoro selesai dengan jumlah skrinsutan:", strconv.Itoa(len(ScreenShotStack)))
	fmt.Println("Tekan Ctrl+C untuk keluar aplikasi")
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	WAclient.Disconnect()

}
