// +build !nacl,!js,cgo,!gofuzz

package oqs_crypto

import (
	"github.com/ethereum/go-ethereum/oqs/oqs_ecdsa"
	"github.com/ethereum/go-ethereum/oqs/oqs_elliptic"
)

// Ecrecover returns the uncompressed public key that created the given signature.
//TODO
func Ecrecover(hash, sig []byte) ([]byte, error) {
	return nil, nil
}

// SigToPub returns the public key that created the given signature.
//TODO
func SigToPub(hash, sig []byte) (*oqs_ecdsa.PublicKey, error) {
	return &oqs_ecdsa.PublicKey{}, nil
}

// Sign calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given digest cannot be chosen by an adversery. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
//TODO
func Sign(digestHash []byte, prv *oqs_ecdsa.PrivateKey) (sig []byte, err error) {
	return nil, nil
}

// VerifySignature checks that the given public key created signature over digest.
// The public key should be in compressed (33 bytes) or uncompressed (65 bytes) format.
// The signature should have the 64 byte [R || S] format.
//TODO
func VerifySignature(pubkey, digestHash, signature []byte) bool {
	return false
}

// DecompressPubkey parses a public key in the 33-byte compressed format.
//TODO
func DecompressPubkey(pubkey []byte) (*oqs_ecdsa.PublicKey, error) {
	return &oqs_ecdsa.PublicKey{}, nil
}

// CompressPubkey encodes a public key to the 33-byte compressed format.
//TODO
func CompressPubkey(pubkey *oqs_ecdsa.PublicKey) []byte {
	return nil
}

// S256 returns an instance of the secp256k1 curve.
//TODO
func S256() oqs_elliptic.Curve {
	return nil
}
