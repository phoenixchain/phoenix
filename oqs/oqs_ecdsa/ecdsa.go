package oqs_ecdsa

import (
	"crypto"
	"github.com/ethereum/go-ethereum/oqs/oqs_elliptic"
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
	PublicKey
	D *big.Int
}

func GenerateKey(c oqs_elliptic.Curve, rand io.Reader) (*PrivateKey, error) {
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
	oqs_elliptic.Curve
	X, Y *big.Int
}

func (pub *PublicKey) Equal(x crypto.PublicKey) bool {
	return false
}
