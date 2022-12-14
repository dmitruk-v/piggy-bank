package entity

import "errors"

var ErrNoOperations = errors.New("no operations found")

type OperationStorage interface {
	GetAll() ([]*CurrencyOperation, error)
	GetLatest() (*CurrencyOperation, error)
	Save(op *CurrencyOperation) error
	DeleteLatest() (*CurrencyOperation, error)
}

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
	Hash       []byte
	PrevHash   []byte
}

func NewCurrencyOperation(optype OperationType, currency Currency, amount float64, providedAt int64, hash []byte, prevHash []byte) *CurrencyOperation {
	return &CurrencyOperation{
		Optype:     optype,
		Currency:   currency,
		Amount:     amount,
		ProvidedAt: providedAt,
		Hash:       hash,
		PrevHash:   prevHash,
	}
}
