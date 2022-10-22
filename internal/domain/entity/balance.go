package entity

import (
	"fmt"
	"sort"
)

type Balance interface {
	HasCurrency(curr Currency) bool
	Add(curr Currency, amount float64) error
	Sub(curr Currency, amount float64) error
	List() BalanceItems
}

type BalanceImpl struct {
	currencies map[Currency]float64
}

type BalanceItem struct {
	Curr   Currency
	Amount float64
}

func NewBalanceImpl(cc ...Currency) *BalanceImpl {
	cmap := make(map[Currency]float64)
	for _, curr := range cc {
		cmap[curr] = 0
	}
	return &BalanceImpl{
		currencies: cmap,
	}
}

func (bal *BalanceImpl) HasCurrency(curr Currency) bool {
	_, ok := bal.currencies[curr]
	return ok
}

func (bal *BalanceImpl) Add(curr Currency, amount float64) error {
	_, ok := bal.currencies[curr]
	if !ok {
		return fmt.Errorf("balance does not have %v currency", curr)
	}
	bal.currencies[curr] += amount
	return nil
}

func (bal *BalanceImpl) Sub(curr Currency, amount float64) error {
	val, ok := bal.currencies[curr]
	if !ok {
		return fmt.Errorf("balance does not have %v currency", curr)
	}
	if val < amount {
		return fmt.Errorf("not enought %v currency: %v, %v needed", curr, val, amount)
	}
	bal.currencies[curr] -= amount
	return nil
}

func (bal *BalanceImpl) List() BalanceItems {
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
