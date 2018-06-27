package core

import (
	"log"
	"strconv"

	funk "github.com/thoas/go-funk"
)

type Wallet interface {
	AvailableBalance() float64
	TotalBalance() float64
	ApplyBlock(block []*Transaction)
	ApplyPendingTransaction(t *Transaction)
	Print()
}

type wallet struct {
	addr       string
	pendingTxs []*Transaction // pending transactions
	utxos      []*TXO         // utxo's addressed to wallet
}

func NewWallet(addr string, pendingTxs []*Transaction, utxos []*TXO) Wallet {
	return &wallet{addr, pendingTxs, utxos}
}

func (w *wallet) Print() {
	for _, utxo := range w.utxos {
		log.Println("UTXO")
		log.Println(utxo.txin.txid)
		log.Println("Address")
		log.Println(utxo.txout.addr)
		log.Println("Change")
		log.Println(strconv.FormatFloat(utxo.txout.c, 'f', -1, 64))
	}
}

// total coins at given address (that don't have inputs in pending txs)
func (w *wallet) AvailableBalance() float64 {
	availableUTXOs := w.availableUTXO()
	return balance(availableUTXOs, w.addr)
}

// total coins at given address in utxos
func (w *wallet) TotalBalance() float64 {
	return balance(w.utxos, w.addr)
}

// return wallet's utxo's that have no txins from pending transactions
func (w *wallet) availableUTXO() []*TXO {
	pendingTxins := []*Txin{}
	for _, pendingTx := range w.pendingTxs {
		pendingTxins = append(pendingTxins, pendingTx.txins...)
	}

	return FilterTXOsWithoutInputs(w.utxos, pendingTxins)
}

// return list of transaction outputs in pending transactions with given address
func (w *wallet) Change() []*Txout {
	pendingTxOuts := ExtractTxouts(w.pendingTxs)
	return funk.Filter(pendingTxOuts, func(txout *Txout) bool {
		return txout.addr == w.addr
	}).([]*Txout)
}

func (w *wallet) ApplyBlock(block []*Transaction) {
	w.utxos = UpdateUTXOsWithBlock(w.addr, w.utxos, block)
	w.pendingTxs = UpdatePendingTxsWithBlock(w.pendingTxs, block)
}

// remove pending transactions that have spent transaction t
func (w *wallet) ApplyPendingTransaction(t *Transaction) {
	w.pendingTxs = UpdatePendingTxsWithBlock(w.pendingTxs, []*Transaction{t})
}

// returns balance of a list of UTXOs for a given address
func balance(utxos []*TXO, addr string) float64 {
	balance := float64(0)
	for _, utxo := range utxos {
		if utxo.txout.addr == addr {
			balance = balance + utxo.txout.c
		}
	}
	return balance
}
