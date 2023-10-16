package main

import (
	"bufio"
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/whatsauth/watoken"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func CheckURLStatus(url string) (status bool, msg string) {
	if !ValidUrl(url) {
		return
	}

	response, err := http.Get(url)
	if err != nil {
		msg = err.Error()
		return
	}
	defer response.Body.Close()
	msg = response.Status
	if msg == "200 OK" {
		if strings.Contains(url, ".github.io") {
			if !strings.Contains(url, "github.com") {
				status = true
			}
		} else if strings.Contains(url, ".google.com") {
			status = true
		}
	}
	return
}

func ValidUrl(urllink string) bool {
	_, er := url.Parse(urllink)
	if er != nil {
		return false
	}
	return true
}

func InputWAGroup() (wag string) {
	beeep.Alert("Pomokit Info", "Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ", "information.png")
	fmt.Println("Please Input Your WhatsApp Group ID with keyword : Iteung minta id grup : ")
	fmt.Scanln(&wag)
	return
}

func InputURLGithub() (hashurl string) {
	var urltask string
	fmt.Println("URL Github Pages atau Google Drive Yang Akan Dikerjakan : ")
	fmt.Scanln(&urltask)
	urlvalid, msgerrurl := CheckURLStatus(urltask)
	for !urlvalid {
		beeep.Alert("Invalid Link", "URL Tidak Valid : "+msgerrurl, "information.png")
		fmt.Println("URL Invalid, Masukkan kembali URL yang benar : ")
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
