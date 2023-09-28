package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/musik"
	"github.com/gen2brain/beeep"
	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func WhatsApp() {
	dbLog := waLog.Stdout("Database", "ERROR", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:wasession.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "ERROR", true)
	WAclient = whatsmeow.NewClient(deviceStore, clientLog)
	//client.AddEventHandler(eventHandler)

	if WAclient.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := WAclient.GetQRChannel(context.Background())
		err = WAclient.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				urlqr := "https://getqr.github.io/#" + evt.Code
				fmt.Println("Membuka URL : ")
				fmt.Println(urlqr)
				openbrowser(urlqr)
				beeep.Alert("Pomokit Info", "Silahkan Scan QR Code Yang Terbuka di Browser dengan Menggunakan Aplikasi WhatsApp", "information.png")
				fmt.Println("Silahkan Scan QR Code Yang Terbuka di Browser dengan Menggunakan Aplikasi WhatsApp")
			} else {
				beeep.Alert("Pomokit Info", "Login WhatsApp:"+evt.Event, "information.png")
				fmt.Println("Login WhatsApp:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = WAclient.Connect()
		if err != nil {
			panic(err)
		}
	}

}

func SendReportTo(filename string, groupid string, hashuserid string) {
	var to = types.JID{
		User:   groupid,
		Server: "g.us",
	}
	filebyte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	//msg := "File dikirim ke server : " + filename
	//atmessage.SendMessage(msg, to, WAclient)
	Hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	msg := "*Pomodoro Report 1 cycle*" + "\nHostname : " + Hostname + "\nIP : https://whatismyipaddress.com/ip/" + strings.TrimSpace(musik.GetIPaddress()) + "\nJumlah ScreenShoot : " + strconv.Itoa(len(ScreenShotStack)) + "\n#" + hashuserid
	atmessage.SendImageMessage(filebyte, msg, to, WAclient)

}

func SendNotifTo(groupid string) {
	var to = types.JID{
		User:   groupid,
		Server: "g.us",
	}
	//msg := "File dikirim ke server : " + filename
	//atmessage.SendMessage(msg, to, WAclient)
	Hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	msg := "*Pomodoro Start 1 cycle*" + "\nHostname : " + Hostname + "\nIP : https://whatismyipaddress.com/ip/" + strings.TrimSpace(musik.GetIPaddress())

	atmessage.SendMessage(msg, to, WAclient)
}
