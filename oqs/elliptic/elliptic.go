package elliptic

import (
	"io"
	"math/big"
)

func GenerateKey(curve Curve, rand io.Reader) (priv []byte, x, y *big.Int, err error) {
	return nil, &big.Int{}, &big.Int{}, nil
}

func Marshal(curve Curve, x, y *big.Int) []byte {
	return nil
}

func MarshalCompressed(curve Curve, x, y *big.Int) []byte {
	return nil
}

func Unmarshal(curve Curve, data []byte) (x, y *big.Int) {
	return &big.Int{}, &big.Int{}
}

func UnmarshalCompressed(curve Curve, data []byte) (x, y *big.Int) {
	return &big.Int{}, &big.Int{}
}

type Curve struct {
}

func P224() Curve {
	return Curve{}
}

func P256() Curve {
	return Curve{}
}

func P384() Curve {
	return Curve{}
}

func P521() Curve {
	return Curve{}
}

type CurveParams struct {
}

func (curve *CurveParams) Add(x1, y1, x2, y2 *big.Int) (*big.Int, *big.Int) {
	return &big.Int{}, &big.Int{}
}

func (curve *CurveParams) Double(x1, y1 *big.Int) (*big.Int, *big.Int) {
	return &big.Int{}, &big.Int{}
}

func (curve *CurveParams) IsOnCurve(x, y *big.Int) bool {
	return false
}

func (curve *CurveParams) Params() *CurveParams {
	return &CurveParams{}
}

func (curve *CurveParams) ScalarBaseMult(k []byte) (*big.Int, *big.Int) {
	return &big.Int{}, &big.Int{}
}

func (curve *CurveParams) ScalarMult(Bx, By *big.Int, k []byte) (*big.Int, *big.Int) {
	return &big.Int{}, &big.Int{}
}
