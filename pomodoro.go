package main

import (
	"bufio"
	"fmt"
	"net/http"
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
	var wag string
	beeep.Alert("Pomokit Info", "Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ", "information.png")
	fmt.Println("Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ")
	fmt.Scanln(&wag)

	var urltask string
	fmt.Println("URL Github Pages Yang Akan Dikerjakan : ")
	fmt.Scanln(&urltask)
	urlvalid, msgerrurl := CheckURLStatus(urltask)
	for !urlvalid {
		beeep.Alert("Invalid Github Pages", "URL Github Pages Tidak Valid : "+msgerrurl, "information.png")
		fmt.Println("URL Github Pages Invalid, Masukkan kembali URL yang benar : ")
		fmt.Scanln(&urltask)
		urlvalid, msgerrurl = CheckURLStatus(urltask)
	}
	hashuserid, err := watoken.EncodeforHours(urltask, PrivateKey, 3)
	if err != nil {
		fmt.Println(err)
	}

	milestone := InputMilestone()
	fmt.Println("Mile Stone", milestone)

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
	SendReportTo(filename, wag, milestone, hashuserid)
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

func CheckURLStatus(url string) (status bool, msg string) {
	response, err := http.Get(url)
	if err != nil {
		msg = err.Error()
	} else {
		msg = response.Status
		if msg == "200 OK" {
			if strings.Contains(url, ".github.io") {
				status = true
			}
		}
	}
	defer response.Body.Close()
	return
}

func InputMilestone() (milestone string) {
	beeep.Alert("Pomokit Info", "Silahkan input rencana yang akan anda kerjakan pada 1 cycle pomodoro sekarang", "information.png")
	fmt.Println("Rencana yang akan anda kerjakan pada 1 cycle pomodoro sekarang : ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		milestone = scanner.Text()
		break

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return

}
