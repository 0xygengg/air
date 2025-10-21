package blockchain

import (
	"crypto/rand"
	"math/big"
	"sync"
)

type Validator struct {
	Address string
	Stake   int
}

var (
	validatorSet = []Validator{}
	stakeLock    sync.Mutex
)

func Stake(address string, amount int) {
	stakeLock.Lock()
	defer stakeLock.Unlock()

	for i := range validatorSet {
		if validatorSet[i].Address == address {
			validatorSet[i].Stake += amount
			return
		}
	}
	validatorSet = append(validatorSet, Validator{
		Address: address,
		Stake:   amount,
	})
}

func Slash(address string, amount int) {
	stakeLock.Lock()
	defer stakeLock.Unlock()

	for i := range validatorSet {
		if validatorSet[i].Address == address {
			validatorSet[i].Stake -= amount
			if validatorSet[i].Stake < 0 {
				validatorSet[i].Stake = 0
			}
			break
		}
	}
}

func SelectValidator() string {
	stakeLock.Lock()
	defer stakeLock.Unlock()

	var total int
	for _, v := range validatorSet {
		total += v.Stake
	}

	if total == 0 {
		return ""
	}

	rnd, _ := rand.Int(rand.Reader, big.NewInt(int64(total)))
	acc := 0
	for _, v := range validatorSet {
		acc += v.Stake
		if int(rnd.Int64()) < acc {
			return v.Address
		}
	}
	return ""
}
