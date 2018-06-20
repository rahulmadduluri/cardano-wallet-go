package main

import (
	"reflect"

	funk "github.com/thoas/go-funk"
)

// UTXO -- unspent transaction outputs
// utxo ∈ UTxO = txin ↦ txout ∈ TxIn ↦ TxOut

type UTXO struct {
	txin  *TransactionInput
	txout *TransactionOutput
}

// UTXOs with given inputs
// ins ◁ utxo = {i ↦ o | i ↦ o ∈ utxo, i ∈ ins}
func FilterUTXOsWithInputs(utxos []*UTXO, txins []*TransactionInput) []*UTXO {
	return funk.Filter(utxos, func(utxo *UTXO) {
		return funk.Contains(txins, utxo.txin) == true
	})
}

// UTXOs without given inputs
// ins ◁/ utxo = {i ↦ o | i ↦ o ∈ utxo, i ∈/ ins}
func FilterUTXOsWithoutInputs(utxos []*UTXO, txins []*TransactionInput) []*UTXO {
	return funk.Filter(utxos, func(utxo *UTXO) {
		return funk.Contains(txins, utxo.txin) == false
	})
}

// UTXOs with given outputs
// utxo ▷ outs = {i ↦ o | i ↦ o ∈ utxo, o ∈ outs}
func FilterUTXOsWithOutputs(utxos []*UTXO, txouts []*TransactionOutput) []*UTXO {
	return funk.Filter(utxos, func(utxo *UTXO) {
		return funk.Contains(txouts, utxo.txout) == true
	})
}
