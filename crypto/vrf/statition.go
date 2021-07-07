package vrf

import (
	"gonum.org/v1/gonum/stat/distuv"
	"math/big"
)

func sortitio(n, p, ratio float64, money uint64) uint64 {
	dist := distuv.Binomial{
		P: p,
		N: n,
	}
	for j := uint64(0); j < money; j++ {
		boundary := dist.CDF(float64(j))
		if ratio <= boundary {
			return j
		}
	}

	return money
}

// Select runs the sortition function and returns the number of time the key was selected
func SelectSort(money uint64, totalMoney uint64, expectedSize float64, vrfOutput []byte) uint64 {
	binomialN := float64(money)
	binomialP := expectedSize / float64(totalMoney)

	t := &big.Int{}
	t.SetBytes(vrfOutput)

	precision := uint(8 * (len(vrfOutput) + 1))
	max, b, err := big.ParseFloat("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 0, precision, big.ToNearestEven)
	if b != 16 || err != nil {
		panic("failed to parse big float constant in sortition")
	}

	h := big.Float{}
	h.SetPrec(precision)
	h.SetInt(t)

	ratio := big.Float{}
	cratio, _ := ratio.Quo(&h, max).Float64()

	return uint64(sortitio(binomialN, binomialP, cratio, money))
}

func VerifySort(money uint64, totalMoney uint64, expectedSize float64, vrfOutput []byte) bool {
	j := SelectSort(money, totalMoney, expectedSize, vrfOutput)
	if j > 0 {
		return true
	}
	return false
}
