package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aiteung/atmessage"
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
				openbrowser("https://getqr.github.io/#" + evt.Code)
				fmt.Println("Silahkan Scan QR Code Yang Terbuka di Browser dengan Menggunakan Aplikasi WhatsApp Linked Device")
			} else {
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

func SendReportTo(filename string, groupid string) {
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

	fmt.Printf("HostName is: %s\n", Hostname)
	atmessage.SendImageMessage(filebyte, "Pomodoro Report", to, WAclient)
}
