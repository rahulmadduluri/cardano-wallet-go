package main

import (
	"core"
)

func main() {
	// reading utxos / pending txs from a file
	addr := "0x123456789"
	wallet := core.NewWallet(addr, []*core.Transaction{}, []*core.TXO{core.TestTXO()})

	// get user input. let users generate a random block or make a specific transaction
	wallet.Print()
}
