package oqs_elliptic

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

type Curve interface {
	// Params returns the parameters for the curve.
	Params() *CurveParams
	// IsOnCurve reports whether the given (x,y) lies on the curve.
	IsOnCurve(x, y *big.Int) bool
	// Add returns the sum of (x1,y1) and (x2,y2)
	Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int)
	// Double returns 2*(x,y)
	Double(x1, y1 *big.Int) (x, y *big.Int)
	// ScalarMult returns k*(Bx,By) where k is a number in big-endian form.
	ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int)
	// ScalarBaseMult returns k*G, where G is the base point of the group
	// and k is an integer in big-endian form.
	ScalarBaseMult(k []byte) (x, y *big.Int)
}

func P224() Curve {
	return nil
}

func P256() Curve {
	return nil
}

func P384() Curve {
	return nil
}

func P521() Curve {
	return nil
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
