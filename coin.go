package main

// HashFunction ...
type HashFunction func(string) string

// ClientBase ...
// The Interface for a coin using HippoTesting Platform.
type ClientBase interface {
	Init(interface{}) error
	Run() error
	SetGood() error
	SetEvil() error
	SetHashFunction(HashFunction) error
	Pause() error
	Stop() error
}

// BlockBase ...
// The Interface for a block using HippoTesting Platform.
type BlockBase interface {
	Init(interface{}) error
	AddTransaction(TransactionBase) (bool, error)
}

// TransactionBase interface ...
// The Interface for a block using HippoTesting Platform.
type TransactionBase interface {
	Init(interface{}) error
}
