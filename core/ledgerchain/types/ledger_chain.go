package types

// LedgerChain defines an interface required for the implementation
// of a linked ledger data structure. A ledger chain is a continous,
// crytographically linked set of ledger with each ledger holding unlimited
// number of cryptographically linked transactions.
type LedgerChain interface {
	Connect(dbAddr string) (interface{}, error)
	Init() error
	GetBackend() string
	CreateLedger(name, cocoonCodeID string, public bool) (interface{}, error)
	Put(ledger, txID, key, value string) (interface{}, error)
	Close() error
}
