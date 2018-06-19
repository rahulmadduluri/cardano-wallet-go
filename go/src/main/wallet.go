package main

import (
	"log"

	funk "github.com/thoas/go-funk"
)

type Wallet interface {
	AvailableBalance() float64
	TotalBalance() float64
	MinimumBalance() float64
	Addresses() []string
	ApplyBlock(b *Block) *wallet
	ApplyPendingTransaction(t *Transaction) *wallet
}

type wallet struct {
	pendingTxs []*Transaction
	utxos      []*UTXO
}

func (w *wallet) AvailableBalance() float64 {
	return 0
}

func (w *wallet) AvailableUTXO() []*UTXO {
	return funk.Filter(w.utxos, func(utxo *UTXO) bool {
		for _, pendingTx := range w.pendingTxs {
			if funk.Contains(pendingTx.TxIns(), utxo.txin) {
				return false
			}
		}
		return true
	})
}

func (w *wallet) Change(addr string) []*TransactionOutput {
	pendingTxOuts := []*TransactionOutput{}
	for _, pendingTx := range w.pendingTxs {
		for _, txout := range pendingTx.TxOuts() {
			if txout.addr == addr {
				append(pendingTxOuts, txout)
			}
		}
	}
	return pendingTxOuts
}

// TODO: impelment
func (w *wallet) TotalBalance() float64 {
	return 0
}

// TODO: impelment
func (w *wallet) MinimumBalance() float64 {
	return 0
}

// TODO: impelment
func (w *wallet) Addresses() []string {
	return w.addr
}

// TODO: impelment
func (w *wallet) ApplyBlock(b *Block) *wallet {
	return wallet{
		[]*Transaction{},
		[]*Transaction{},
	}
}

// TODO: impelment
func (w *wallet) ApplyPendingTransaction(t *Transaction) *wallet {
	return wallet{
		[]*Transaction{},
		[]*Transaction{},
	}
}
