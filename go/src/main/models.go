package main

import ()

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

// UTXO -- unspent transaction outputs
// utxo ∈ UTxO = txin |→ txout ∈ TxIn |→ TxOut

type UTXO struct {
	txin  *TransactionInput
	txout *TransactionOutput
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
	UTXOs() []*UTXO
	TxIns() []*TransactionInput
	TxOuts() []*TransactionOutput
}

type transaction struct {
	txins  []*TransactionInput
	txouts []*TransactionOutput
}

func (t *transaction) TxId() int {
	// TODO: implement
	return 0
}

func (t *transaction) UTXOs() []*UTXO {
	return
}

func (t *transaction) TxIns() []*TransactionInput {
	return t.txins
}

func (t *transaction) TxOuts() []*TransactionOutput {
	return t.txouts
}
