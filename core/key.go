package core

type Key struct {
}

type PrivateKey struct {
	Key
}

type PublicKey struct {
	Key
}

func NewPrivateKey() *PrivateKey {
	pk := &PrivateKey{}

	return pk
}

func NewPublicKey() *PublicKey {
	pk := &PublicKey{}

	return pk
}

func (k *Key) Validate() bool {

	return true
}