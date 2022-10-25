package entity

type Balance interface {
	Amount(curr Currency) float64
	Add(curr Currency, amount float64) error
	Sub(curr Currency, amount float64) error
	List() BalanceItems
}
