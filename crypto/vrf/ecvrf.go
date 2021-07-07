package vrf

import (
	"crypto/ecdsa"
	"github.com/vechain/go-ecvrf"
)
var (
	secp256Vrf ecvrf.VRF
)

func init()  {
	secp256Vrf =ecvrf.NewSecp256k1Sha256Tai()
}

func Prove(sk *ecdsa.PrivateKey, alpha []byte) (beta, pi []byte, err error){
	return secp256Vrf.Prove(sk,alpha)
}

func Verify(pk *ecdsa.PublicKey, alpha, pi []byte) (beta []byte, err error){
	return secp256Vrf.Verify(pk,alpha,pi)
}