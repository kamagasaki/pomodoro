package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
	"github.com/whatsauth/watoken"
)

func CheckURLStatus(url string) (status bool, msg string) {
	response, err := http.Get(url)
	if err != nil {
		msg = err.Error()
	} else {
		msg = response.Status
		if msg == "200 OK" {
			if strings.Contains(url, ".github.io") {
				if !strings.Contains(url, "github.com") {
					status = true
				}

			}
		}
	}
	defer response.Body.Close()
	return
}

func InputWAGroup() (wag string) {
	beeep.Alert("Pomokit Info", "Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ", "information.png")
	fmt.Println("Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ")
	fmt.Scanln(&wag)
	return
}

func InputURLGithub() (hashurl string) {
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
	hashurl, err := watoken.EncodeforHours(urltask, PrivateKey, 3)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InputMilestone() (milestone string) {
	beeep.Alert("Pomokit Info", "Silahkan input rencana yang akan anda kerjakan pada 1 cycle pomodoro sekarang", "information.png")
	fmt.Println("Rencana yang akan anda kerjakan pada 1 cycle pomodoro sekarang : ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		milestone = scanner.Text()
		if len(milestone) > 17 {
			break
		} else {
			beeep.Alert("Pomokit Info", "Rencana belum diisi atau terlalu pendek", "information.png")
			fmt.Println("Rencana belum diisi atau terlalu pendek, Rencana Anda : ")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return

}
