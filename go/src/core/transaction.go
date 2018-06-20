package main

import (
	"reflect"
)

// TransactionInput
// txin ∈ TxIn = (txid, ix) ∈ TxId × Ix

type TransactionInput struct {
	txid string
	ix   int
}

// TransactionOutput
// txout ∈ TxOut = (addr, c) ∈ Addr × Coin

type TransactionOutput struct {
	addr string
	c    float64
}

// Block
// b ∈ Block = tx ∈ P(Tx)

type Block struct {
	txs []*Transaction
}

// Transaction
// tx ∈ Tx = (inputs, outputs) ∈ P(TxIn) × (Ix |→ TxOut)

type Transaction interface {
	TxId() int
	TxIns() []*TransactionInput
	TxOuts() []*TransactionOutput
	HasTxIn(txin *TransactionInput) bool
	HasTxOut(txout *TransactionOutput) bool
}

type transaction struct {
	txins  []*TransactionInput
	txouts []*TransactionOutput
}

func (t *transaction) TxId() int {
	// TODO: implement
	return 0
}

func (t *transaction) TxIns() []*TransactionInput {
	return t.txins
}

func (t *transaction) TxOuts() []*TransactionOutput {
	return t.txouts
}

func (t *transaction) HasTxIn(txin *TransactionInput) bool {
	for _, txinX := range t.txins {
		if reflect.DeepEqual(txinX, txin) {
			return true
		}
	}
	return false
}

func (t *transaction) HasTxOut(txout *TransactionOutput) bool {
	for _, txoutX := range t.txouts {
		if reflect.DeepEqual(txoutX, txout) {
			return true
		}
	}
	return false
}

// Utility functions

func TransactionsToTxOuts(txs []*Transaction) []*TransactionOutput {
	txouts := []*core.TransactionOutput{}
	for _, tx := range txs {
		append(txouts, tx.TxOuts()...)
	}
	return txouts
}
