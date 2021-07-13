package vrf

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSortitionBasic(t *testing.T) {
	hitcount := uint64(0)
	const N = 1000
	const expectedSize = 10
	const myMoney = 100
	const totalMoney = 200
	for i := 0; i < N; i++ {
		var vrfOutput [32]byte
		rand.Read(vrfOutput[:])
		selected := SelectSort(myMoney, totalMoney, expectedSize, vrfOutput[:])
		fmt.Printf("i=%d, vrfOutput=%x, selected=%d\n", i, vrfOutput, selected)
		hitcount += selected
	}
	expected := uint64(N * expectedSize / 2)
	var d uint64
	if expected > hitcount {
		d = expected - hitcount
	} else {
		d = hitcount - expected
	}
	// within 2% good enough
	maxd := expected / 50
	fmt.Printf("expected=%d, hitcount=%d, d=%d, maxd=%d\n", expected, hitcount, d, maxd)

	if d > maxd {
		fmt.Errorf("wanted %d selections but got %d, d=%d, maxd=%d", expected, hitcount, d, maxd)
	}
}

type Node struct {
	beta   []byte
	money  uint64
	hit    uint64
	rate   float64
	priKey *ecdsa.PrivateKey
}

func TestSortitionSim(t *testing.T) {
	const N = 1000
	const expectedSize = 3
	const NodeSize = 50
	nodes := make([]*Node, NodeSize)
	var totalMoney uint64
	for i := 0; i < NodeSize; i++ {
		money := uint64(rand.Int63n(40000000))
		sk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		nodes[i] = &Node{
			money:  money,
			priKey: sk,
		}
		totalMoney += money
	}
	totalCount := 0
	var seed [32]byte

	fmt.Printf("N=%d, total_money=%d, expectedSize=%d\n", N, totalMoney, expectedSize)
	for j := 0; j < N; j++ {
		rand.Read(seed[:])
		for i := 0; i < NodeSize; i++ {
			beta, _, err := Prove(nodes[i].priKey, seed[:])
			if err != nil {
				t.Fatal(err)
			}
			selected := SelectSort(nodes[i].money, totalMoney, expectedSize, beta)
			if selected > 0 {
				nodes[i].hit += 1
				totalCount++
			}
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].money > nodes[j].money
	})
	fmt.Printf("%10s%10s\n", "money", "selected", "rate")
	for i := 0; i < NodeSize; i++ {
		nodes[i].rate = float64(nodes[i].money) / float64(totalMoney) * 100
		fmt.Printf("%10d%10d%10f\n", nodes[i].money, nodes[i].hit, nodes[i].rate)
	}
	fmt.Printf("totalMoney=%d,hitCount=%d,totalCount=%d,hitRate=%f\n", totalMoney, totalCount, len(nodes)*N, float64(totalCount)/float64(len(nodes)*N))
	//fmt.Printf("End\n")
}

func TestSelectAlgoSpeed(t *testing.T) {
	const N = 10000
	const expectedSize = 10
	const myMoney = 100
	const totalMoney = 200
	var count uint64
	start := time.Now()
	for i := 0; i < N; i++ {
		var seed [32]byte
		rand.Read(seed[:])

		sk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		alpha := seed[:]

		beta, _, err := Prove(sk, alpha)

		selected := SelectSort(myMoney, totalMoney, expectedSize, beta)
		if selected > 0 {
			count += 1
		}
	}
	dis := time.Since(start)
	fmt.Printf("time=%s\n", dis)
}

func TestSelectAlgoSpeed2(t *testing.T) {
	const expectedSize = 1
	const myMoney = 1
	const totalMoney = 1

	seed, err := hex.DecodeString("a71c2340636002424a086eae6e8b319ef15aebd0bc71dd0773d84de0ea2715b4")
	if err != nil {
		panic(err)
	}

	sk, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	alpha := seed[:]

	beta, _, err := Prove(sk, alpha)

	selected := SelectSort(myMoney, totalMoney, expectedSize, beta)
	fmt.Printf("time=%d\n", selected)
}

func TestSortitionNode(t *testing.T) {
	SortitionNode(t)
}

func SortitionNode(t *testing.T) uint64 {
	var alpha [32]byte
	rand.Read(alpha[:])

	const N = 50
	const expectedSize = 3
	totalMoney := uint64(0)
	hitcount := uint64(0)
	hitCnt := uint64(0)

	nodes := make([]*Node, N)
	for i := 0; i < N; i++ {
		money := uint64(rand.Int63n(40000000))
		//money := uint64(rand.Int63n(1000000000000000))
		//money := uint64(rand.Int63n(math.MaxInt64 / N))
		sk, err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		beta, _, err := Prove(sk, alpha[:])

		nodes[i] = &Node{
			beta:  beta,
			money: money,
		}
		totalMoney += money
		if money > totalMoney {
			panic("money too large")
		}
	}

	for i := 0; i < N; i++ {
		selected := SelectSort(nodes[i].money, totalMoney, expectedSize, nodes[i].beta)
		//fmt.Printf("i=%d, beta=%x, myMoney=%d, selected=%d\n", i, nodes[i].beta, nodes[i].money, selected)
		hitcount += selected
		if selected > 0 {
			hitCnt++
		}
	}
	//fmt.Printf("hitcount=%d, hitCnt=%d, totalMoney=%d\n", hitcount, hitCnt, totalMoney)
	return hitCnt
}

func TestSortitionMoreNode(t *testing.T) {
	totalNode := uint64(0)
	num := uint64(1000)
	for i := uint64(0); i < num; i++ {
		totalNode += SortitionNode(t)
	}
	fmt.Println(float64(totalNode) / float64(num))
}
