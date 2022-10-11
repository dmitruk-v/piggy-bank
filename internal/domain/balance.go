package domain

import (
	"fmt"
)

type Balance struct {
	currencies map[Currency]float64
}

func NewBalance(cc ...Currency) *Balance {
	cm := make(map[Currency]float64)
	for i := range cc {
		curr := cc[i]
		cm[curr] = 0
	}
	return &Balance{
		currencies: cm,
	}
}

func (bal *Balance) Add(currency Currency, amount float64) error {
	_, ok := bal.currencies[currency]
	if !ok {
		return fmt.Errorf("balance does not have %v currency", currency)
	}
	bal.currencies[currency] += amount
	return nil
}

func (bal *Balance) Sub(currency Currency, amount float64) error {
	val, ok := bal.currencies[currency]
	if !ok {
		return fmt.Errorf("balance does not have %v currency", currency)
	}
	if val < amount {
		return fmt.Errorf("not enought %v currency: %v, %v needed", currency, val, amount)
	}
	bal.currencies[currency] -= amount
	return nil
}
