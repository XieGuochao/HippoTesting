package hippotesting

import "crypto"

// HashFunction ...
type HashFunction func(string) string

// KeyPair ...
type KeyPair struct {
	PublicKey  crypto.PublicKey
	PrivateKey crypto.PrivateKey
}

// NewKeyPair ...
func NewKeyPair(publicKey, privateKey []byte) *KeyPair {
	kp := new(KeyPair)
	kp.PrivateKey = privateKey
	kp.PublicKey = publicKey
	return kp
}

// ClientBase ...
// The Interface for a coin using HippoTesting Platform.
type ClientBase interface {
	Init(interface{}) error
	Run() error
	SetGood() error
	SetEvil() error
	SetHashFunction(HashFunction) error
	GetHashFunction() HashFunction
	Pause() error
	Stop() error

	GetGenisis() *BlockBase

	GetKeyPair() *KeyPair

	Sync() error
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
