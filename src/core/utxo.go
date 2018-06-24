package core

import (
	"reflect"

	"github.com/jinzhu/copier"
	funk "github.com/thoas/go-funk"
)

// TXO -- transaction outputs (can be signed or unsigned)
// utxo ∈ UTxO = txin ↦ txout ∈ TxIn ↦ TxOut

type TXO struct {
	txin  *Txin
	txout *Txout
}

func TestTXO() *TXO {
	return &TXO{
		txin: &Txin{
			txid: "TestTXID",
			ix:   0,
		},
		txout: &Txout{
			addr: "0x23457890",
			c:    8.0,
		},
	}
}

// TXOs with given inputs
// ins ◁ utxo = {i ↦ o | i ↦ o ∈ utxo, i ∈ ins}
func FilterTXOsWithInputs(txos []*TXO, txins []*Txin) []*TXO {
	filteredTXOs := funk.Filter(txos, func(txo *TXO) bool {
		for _, txin := range txins {
			if reflect.DeepEqual(txo.txin, txin) {
				return true
			}
		}
		return false
	}).([]*TXO)
	return filteredTXOs
}

// TXOs without given inputs
// ins ◁/ utxo = {i ↦ o | i ↦ o ∈ utxo, i ∈/ ins}
func FilterTXOsWithoutInputs(txos []*TXO, txins []*Txin) []*TXO {
	filteredTXOs := funk.Filter(txos, func(txo *TXO) bool {
		for _, txin := range txins {
			if reflect.DeepEqual(txo.txin, txin) {
				return true
			}
		}
		return false
	}).([]*TXO)
	return filteredTXOs
}

// TXOs with given outputs
// utxo ▷ outs = {i ↦ o | i ↦ o ∈ utxo, o ∈ outs}
func FilterTXOsWithOutputs(txos []*TXO, txouts []*Txout) []*TXO {
	filteredTXOs := funk.Filter(txos, func(txo *TXO) bool {
		return funk.Contains(txouts, txo.txout) == true
	}).([]*TXO)
	return filteredTXOs
}

// new b = txins b ◁/ (txouts b ▷ TxOutours)
func UpdateUTXOsWithBlock(addr string, utxos []*TXO, block []*Transaction) []*TXO {
	// 1. Get all NEW utxos from block that are addressed to me
	// 2. Add these outputs to our utxo list
	// 3. Remove the UTXOs where the transaction inputs are SPENT (used) in the block
	allTXOs := []*TXO{}
	copier.Copy(utxos, allTXOs)
	allTXOs = append(allTXOs, ExtractTXOs(block)...)
	updatedUTXOs := funk.Filter(allTXOs, func(txo *TXO) bool {
		isOutputSpent := false
		for _, tx := range block {
			if tx.txid == txo.txin.txid {
				isOutputSpent = true
			}
		}
		// ensure output is not spent && is addressed to me
		return isOutputSpent == false && txo.txout.addr == addr
	}).([]*TXO)
	return updatedUTXOs
}

func ExtractTXOs(block []*Transaction) []*TXO {
	txos := []*TXO{}
	for _, tx := range block {
		for i, txout := range tx.txouts {
			txo := &TXO{
				txin: &Txin{
					txid: tx.txid,
					ix:   i,
				},
				txout: &Txout{
					addr: txout.addr,
					c:    txout.c,
				},
			}
			txos = append(txos, txo)
		}
	}
	return txos
}
