package domain

type OperationType int64

const (
	DepositOperation OperationType = iota + 1
	WithdrawOperation
)

type CurrencyOperation struct {
	Optype     OperationType
	Currency   Currency
	Amount     float64
	ProvidedAt int64
}

func NewCurrencyOperation(optype OperationType, currency Currency, amount float64, providedAt int64) *CurrencyOperation {
	return &CurrencyOperation{
		Optype:     optype,
		Currency:   currency,
		Amount:     amount,
		ProvidedAt: providedAt,
	}
}
