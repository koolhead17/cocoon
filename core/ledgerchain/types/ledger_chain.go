package types

// LedgerChain defines an interface required for the implementation
// of a linked ledger data structure. A ledger chain is a continous,
// crytographically linked set of ledger with each ledger holding unlimited
// number of cryptographically linked transactions.
type LedgerChain interface {
	Connect(dbAddr string) (interface{}, error)
	Init() error
	GetBackend() string
	CreateLedger(name, cocoonCodeID string, public bool) (*Ledger, error)
	GetLedger(name string) (*Ledger, error)
	ListLedgers(cocoonCodeID string) ([]*Ledger, error)
	Put(txID, key, value string) (*Transaction, error)
	Get(key string) (*Transaction, error)
	GetByID(txID string) (*Transaction, error)
	Close() error
}
