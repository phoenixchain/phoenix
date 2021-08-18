package cryptor

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/oqs/ecdsa"
	"github.com/ethereum/go-ethereum/oqs/elliptic"
	"math/big"
)

//SignatureLength indicates the byte length required to carry a signature with recovery id.
const SignatureLength = 64 + 1 // 64 bytes ECDSA signature + 1 byte recovery id

// RecoveryIDOffset points to the byte offset within the signature that contains the recovery id.
const RecoveryIDOffset = 64

// DigestLength sets the signature digest exact length
const DigestLength = 32

func CompressPubkey(pubkey *ecdsa.PublicKey) []byte {
	return nil
}

func CreateAddress(b common.Address, nonce uint64) common.Address {
	return [20]byte{}
}

func CreateAddress2(b common.Address, salt [32]byte, inithash []byte) common.Address {
	return [20]byte{}
}

func DecompressPubkey(pubkey []byte) (*ecdsa.PublicKey, error) {
	return &ecdsa.PublicKey{}, nil
}

func Ecrecover(hash, sig []byte) ([]byte, error) {
	return nil, nil
}

func FromECDSA(priv *ecdsa.PrivateKey) []byte {
	return nil
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	return nil
}

func GenerateKey() (*ecdsa.PrivateKey, error) {
	return &ecdsa.PrivateKey{}, nil
}

func HashData(kh KeccakState, data []byte) (h common.Hash) {
	return [32]byte{}
}

func HexToECDSA(hexkey string) (*ecdsa.PrivateKey, error) {
	return &ecdsa.PrivateKey{}, nil
}

func Keccak256(data ...[]byte) []byte {
	return nil
}

func Keccak256Hash(data ...[]byte) (h common.Hash) {
	return [32]byte{}
}

func Keccak512(data ...[]byte) []byte {
	return nil
}

func LoadECDSA(file string) (*ecdsa.PrivateKey, error) {
	return &ecdsa.PrivateKey{}, nil
}

func PubkeyToAddress(p ecdsa.PublicKey) common.Address {
	return [20]byte{}
}

func S256() elliptic.Curve {
	return elliptic.Curve{}
}

func SaveECDSA(file string, key *ecdsa.PrivateKey) error {
	return nil
}

func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	return &ecdsa.PublicKey{}, nil
}

func Sign(digestHash []byte, prv *ecdsa.PrivateKey) (sig []byte, err error) {
	return nil, nil
}

func ToECDSA(d []byte) (*ecdsa.PrivateKey, error) {
	return &ecdsa.PrivateKey{}, nil
}

func ToECDSAUnsafe(d []byte) *ecdsa.PrivateKey {
	return &ecdsa.PrivateKey{}
}

func UnmarshalPubkey(pub []byte) (*ecdsa.PublicKey, error) {
	return &ecdsa.PublicKey{}, nil
}

func ValidateSignatureValues(v byte, r, s *big.Int, homestead bool) bool {
	return false
}

func VerifySignature(pubkey, digestHash, signature []byte) bool {
	return false
}

type KeccakState struct {
}

func NewKeccakState() KeccakState {
	return KeccakState{}
}
