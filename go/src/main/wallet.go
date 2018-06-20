package main

import (
	"log"

	"core"

	funk "github.com/thoas/go-funk"
)

type Wallet interface {
	AvailableBalance() float64
	TotalBalance() float64
	MinimumBalance() float64
	ApplyBlock(b *core.Block) *wallet
	ApplyPendingTransaction(t *core.Transaction) *wallet
}

type wallet struct {
	pendingTxs []*core.Transaction // pending transactions
	utxos      []*core.UTXO        // utxo's addressed to wallet
}

// total coins at given address (that don't have inputs in pending txs)
func (w *wallet) AvailableBalance(addr string) float64 {
	availableUTXOs := w.availableUTXO()
	return balance(availableUTXOs, addr)
}

// total coins at given address in utxos
func (w *wallet) TotalBalance(addr string) float64 {
	return balance(w.utxos, addr)
}

func (w *wallet) MinimumBalance() float64 {
	return 0
}

// return wallet's utxo's that have no txins from pending transactions
func (w *wallet) availableUTXO() []*core.UTXO {
	pendingTxIns := []*core.TransactionInput{}
	for _, pendingTx := range w.pendingTxs {
		append(pendingTxIns, pendingTx.TxIns()...)
	}

	return core.FilterUTXOsWithoutInputs(w.utxos, pendingTxIns)
}

// return list of transaction outputs in pending transactions with given address
func (w *wallet) Change(addr string) []*core.TransactionOutput {
	pendingTxOuts := core.TransactionToTxOuts(w.pendingTxs)
	return funk.Filter(pendingTxOuts, func(txout *core.TransactionOutput) {
		return txout.addr == addr
	})
}

func (w *wallet) ApplyBlock(b *core.Block) *wallet {
	// 1. update UTXOs
	// 2. update pending
	return wallet{
		[]*core.Transaction{},
		[]*core.Transaction{},
	}
}

// remove pending transactions that have spent transaction t
func (w *wallet) ApplyPendingTransaction(t *core.Transaction) *wallet {
	return wallet{
		[]*core.Transaction{},
		[]*core.Transaction{},
	}
}

// returns balance of a list of UTXOs for a given address
func balance(utxos []*core.UTXO, addr string) float64 {
	balance := 0.0
	for _, utxo := range utxos {
		if utxo.txout.addr == addr {
			balance = balance + utxo.txout.c
		}
	}
	return balance
}

func updateUTXOs(utxos []*core.UTXO, b core.Block) *core.UTXO {

}
