package vrf

import (
	"crypto/ecdsa"
	"github.com/vechain/go-ecvrf"
	"testing"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestVrf(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// the input to be hashed by the VRF
	alpha := "Hello VeChain"

	// `beta`: the VRF hash output
	// `pi`: the VRF proof
	_, pi, err := ecvrf.NewSecp256k1Sha256Tai().Prove(privateKey, []byte(alpha))
	if err != nil {
		// something wrong.
		// most likely sk is not properly loaded.
		t.Fatal(err)
	}
	// `pi` is the VRF proof
	_, err = ecvrf.NewSecp256k1Sha256Tai().Verify(publicKeyECDSA, []byte(alpha), pi)
	if err != nil {
		// invalid proof
		t.Fatal(err)
	}
}