package ecdsa

import (
	"crypto"
	"github.com/ethereum/go-ethereum/oqs/elliptic"
	"io"
	"math/big"
)

func Sign(rand io.Reader, priv *PrivateKey, hash []byte) (r, s *big.Int, err error) {
	return &big.Int{}, &big.Int{}, nil
}

func SignASN1(rand io.Reader, priv *PrivateKey, hash []byte) ([]byte, error) {
	return nil, nil
}

func Verify(pub *PublicKey, hash []byte, r, s *big.Int) bool {
	return false
}

func VerifyASN1(pub *PublicKey, hash, sig []byte) bool {
	return false
}

type PrivateKey struct {
}

func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) {
	return &PrivateKey{}, nil
}

func (priv *PrivateKey) Equal(x crypto.PrivateKey) bool {
	return false
}

func (priv *PrivateKey) Public() crypto.PublicKey {
	return nil
}

func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	return nil, nil
}

type PublicKey struct {
}

func (pub *PublicKey) Equal(x crypto.PublicKey) bool {
	return false
}
