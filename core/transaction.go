package core

import (
	"time"
	"hash"
)

type Transaction struct {
	From      Address
	To        Address
	StartTime time.Time
	EndTime   time.Time
	Payload   Payload
	Contract  Contract
	Receipt   Receipt
	Hash      hash.Hash
}

func NewTransaction() *Transaction {
	t := &Transaction{}

	return t
}

func (t *Transaction) Validate() bool {

	return true
}