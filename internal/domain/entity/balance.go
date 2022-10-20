package entity

import (
	"fmt"
	"sort"
)

type Balance struct {
	currencies map[Currency]float64
}

type BalanceItem struct {
	Curr   Currency
	Amount float64
}

func NewBalance(cc []Currency) *Balance {
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

func (bal *Balance) List() BalanceItems {
	var list BalanceItems
	for curr, amt := range bal.currencies {
		list = append(list, BalanceItem{Curr: curr, Amount: amt})
	}
	sort.Sort(list)
	return list
}

type BalanceItems []BalanceItem

func (s BalanceItems) Len() int           { return len(s) }
func (s BalanceItems) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BalanceItems) Less(i, j int) bool { return s[i].Curr < s[j].Curr }
