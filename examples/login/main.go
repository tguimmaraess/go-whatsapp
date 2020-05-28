package main

import (
	"fmt"
	"github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"os"
	"time"
)

func main() {
	wac, err := whatsapp.NewConn(5 * time.Second)
	//try uncomenting one of these lines if tou can't get it working
	// wac.SetClientVersion(2, 2021, 3)
	// wac.SetClientVersion(2, 2021, 4)
	// wac.SetClientVersion(2, 2021, 2)



	
	if err != nil {
		panic(err)
	}

	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	session, err := wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
	}
	fmt.Printf("login successful, session: %v\n", session)
}
