package core

import (
	"reflect"
)

// TransactionInput
// txin ∈ TxIn = (txid, ix) ∈ TxId × Ix

type Txin struct {
	txid string // the transaction ID hash input is from
	ix   int    // the index of the output of the transaction input is from
}

// TransactionOutput
// txout ∈ TxOut = (addr, c) ∈ Addr × Coin

type Txout struct {
	addr string
	c    float64
}

// Transaction
// tx ∈ Tx = (inputs, outputs) ∈ P(TxIn) × (Ix ↦ TxOut)

type Transaction struct {
	txid   string
	txins  []*Txin
	txouts []*Txout
}

// Utility functions

func HasTxIn(t *Transaction, txin *Txin) bool {
	for _, txinX := range t.txins {
		if reflect.DeepEqual(txinX, txin) {
			return true
		}
	}
	return false
}

func HasTxOut(t *Transaction, txout *Txout) bool {
	for _, txoutX := range t.txouts {
		if reflect.DeepEqual(txoutX, txout) {
			return true
		}
	}
	return false
}

func ExtractTxins(txs []*Transaction) []*Txin {
	txins := []*Txin{}
	for _, tx := range txs {
		txins = append(txins, tx.txins...)
	}
	return txins
}

func ExtractTxouts(txs []*Transaction) []*Txout {
	txouts := []*Txout{}
	for _, tx := range txs {
		txouts = append(txouts, tx.txouts...)
	}
	return txouts
}

func UpdatePendingTxsWithBlock(pendingTxs []*Transaction, block []*Transaction) []*Transaction {
	return pendingTxs
}
