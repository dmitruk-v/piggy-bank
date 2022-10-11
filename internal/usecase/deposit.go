package usecase

import (
	"fmt"
	"time"

	"github.com/dmitruk-v/piggy-bank/internal/domain"
)

type OperationStorage interface {
	List() ([]*domain.CurrencyOperation, error)
	Save(op *domain.CurrencyOperation) error
	DeleteLast() (*domain.CurrencyOperation, error)
}

type DepositUseCase struct {
	balance   *domain.Balance
	opStorage OperationStorage
}

func NewDepositUseCase(balance *domain.Balance, opStorage OperationStorage) *DepositUseCase {
	return &DepositUseCase{
		balance:   balance,
		opStorage: opStorage,
	}
}

func (ucase *DepositUseCase) Execute(currency domain.Currency, amount float64) error {
	if err := ucase.balance.Add(currency, amount); err != nil {
		return fmt.Errorf("execute deposit operation: %v", err)
	}
	op := domain.NewCurrencyOperation(domain.DepositOperation, currency, amount, time.Now().Unix())
	if err := ucase.opStorage.Save(op); err != nil {
		return fmt.Errorf("execute deposit operation: %v", err)
	}
	return nil
}
