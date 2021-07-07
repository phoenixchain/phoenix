package vrf

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"github.com/ethereum/go-ethereum/crypto"
)

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

type User struct {
	name           string
	money          uint64
	selected_count int
}

func TestSortitionSim(t *testing.T) {
	const N = 1000
	const expectedSize = 5
	users := []User{
		{"A", 100, 0},
		{"B", 200, 0},
		{"C", 300, 0},
		{"D", 400, 0},
		{"E", 500, 0},
		{"A1", 1000, 0},
		{"A2", 10000, 0},
		{"A3", 100000, 0},
		//{"A4",1000000, 0},
		//{"A5",10000000, 0},
	}
	var totalMoney uint64
	for i := 0; i < len(users); i++ {
		totalMoney += users[i].money
	}

	fmt.Printf("N=%d, total_money=%d, expectedSize=%d\n", N, totalMoney, expectedSize)
	fmt.Printf("%10s%10s%10s\n", "name", "money", "selected")
	for i := 0; i < len(users); i++ {
		for j := 0; j < N; j++ {
			var vrfOutput [32]byte
			rand.Read(vrfOutput[:])
			selected := SelectSort(users[i].money, totalMoney, expectedSize, vrfOutput[:])
			if selected > 0 {
				users[i].selected_count += 1
			}
		}
		fmt.Printf("%10s%10d%10d\n", users[i].name, users[i].money, users[i].selected_count)
	}
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

		sk,err := crypto.GenerateKey()
		if err != nil {
			t.Fatal(err)
		}
		alpha := seed[:]

		beta, _, err := Prove(sk,alpha)

		selected := SelectSort(myMoney, totalMoney, expectedSize, beta)
		if selected > 0 {
			count += 1
		}
	}
	dis := time.Since(start)
	fmt.Printf("time=%s\n", dis)
}
