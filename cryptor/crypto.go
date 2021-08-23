package cryptor

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/oqs/oqs_ecdsa"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
	"hash"
	"math/big"
)

//SignatureLength indicates the byte length required to carry a signature with recovery id.
const SignatureLength = 64 + 1 // 64 bytes ECDSA signature + 1 byte recovery id

// RecoveryIDOffset points to the byte offset within the signature that contains the recovery id.
const RecoveryIDOffset = 64

// DigestLength sets the signature digest exact length
const DigestLength = 32

// CreateAddress creates an ethereum address given the bytes and the nonce
func CreateAddress(b common.Address, nonce uint64) common.Address {
	data, _ := rlp.EncodeToBytes([]interface{}{b, nonce})
	return common.BytesToAddress(Keccak256(data)[12:])
}

// CreateAddress2 creates an ethereum address given the address bytes, initial
// contract code hash and a salt.
func CreateAddress2(b common.Address, salt [32]byte, inithash []byte) common.Address {
	return common.BytesToAddress(Keccak256([]byte{0xff}, b.Bytes(), salt[:], inithash)[12:])
}

func FromECDSA(priv *oqs_ecdsa.PrivateKey) []byte {
	return nil
}

//TODO
func FromECDSAPub(pub *oqs_ecdsa.PublicKey) []byte {
	return nil
}

//TODO
func GenerateKey() (*oqs_ecdsa.PrivateKey, error) {
	return &oqs_ecdsa.PrivateKey{}, nil
}

// HashData hashes the provided data using the KeccakState and returns a 32 byte hash
func HashData(kh KeccakState, data []byte) (h common.Hash) {
	kh.Reset()
	kh.Write(data)
	kh.Read(h[:])
	return h
}

//TODO
func HexToECDSA(hexkey string) (*oqs_ecdsa.PrivateKey, error) {
	return &oqs_ecdsa.PrivateKey{}, nil
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := NewKeccakState()
	for _, b := range data {
		d.Write(b)
	}
	d.Read(b)
	return b
}

// Keccak256Hash calculates and returns the Keccak256 hash of the input data,
// converting it to an internal Hash data structure.
func Keccak256Hash(data ...[]byte) (h common.Hash) {
	d := NewKeccakState()
	for _, b := range data {
		d.Write(b)
	}
	d.Read(h[:])
	return h
}

// Keccak512 calculates and returns the Keccak512 hash of the input data.
func Keccak512(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak512()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

//TODO
func LoadECDSA(file string) (*oqs_ecdsa.PrivateKey, error) {
	return &oqs_ecdsa.PrivateKey{}, nil
}

func PubkeyToAddress(p oqs_ecdsa.PublicKey) common.Address {
	return [20]byte{}
}

//TODO
func SaveECDSA(file string, key *oqs_ecdsa.PrivateKey) error {
	return nil
}

//TODO
func ToECDSA(d []byte) (*oqs_ecdsa.PrivateKey, error) {
	return &oqs_ecdsa.PrivateKey{}, nil
}

// ToECDSAUnsafe blindly converts a binary blob to a private key. It should almost
// never be used unless you are sure the input is valid and want to avoid hitting
// errors due to bad origin encoding (0 prefixes cut off).
//TODO
func ToECDSAUnsafe(d []byte) *oqs_ecdsa.PrivateKey {
	return &oqs_ecdsa.PrivateKey{}
}

// UnmarshalPubkey converts bytes to a secp256k1 public key.
//TODO
func UnmarshalPubkey(pub []byte) (*oqs_ecdsa.PublicKey, error) {
	return &oqs_ecdsa.PublicKey{}, nil
}

func ValidateSignatureValues(v byte, r, s *big.Int, homestead bool) bool {
	return false
}

// KeccakState wraps sha3.state. In addition to the usual hash methods, it also supports
// Read to get a variable amount of data from the hash state. Read is faster than Sum
// because it doesn't copy the internal state, but also modifies the internal state.
type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

func NewKeccakState() KeccakState {
	return sha3.NewLegacyKeccak256().(KeccakState)
}
